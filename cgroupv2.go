package clistat

import (
	"errors"
	"io/fs"
	"path/filepath"
	"strconv"

	"github.com/spf13/afero"
	"golang.org/x/xerrors"
	"tailscale.com/types/ptr"
)

// Paths for CgroupV2.
// Ref: https://docs.kernel.org/6.17/admin-guide/cgroup-v2.html
const (
	// Contains quota and period in microseconds separated by a space.
	cgroupV2CPUMax = "cpu.max"
	// Contains current CPU usage under usage_usec
	cgroupV2CPUStat = "cpu.stat"
	// Contains current cgroup memory usage in bytes.
	cgroupV2MemoryUsageBytes = "memory.current"
	// Contains max cgroup memory usage in bytes.
	cgroupV2MemoryMaxBytes = "memory.max"
	// Other memory stats - we are interested in total_inactive_file
	cgroupV2MemoryStat = "memory.stat"

	// Default period for cpu.max as documented in the kernel docs.
	// The default period is 100000 microseconds (100ms).
	// Ref: https://docs.kernel.org/6.17/admin-guide/cgroup-v2.html#cpu-interface-files
	cgroupV2DefaultPeriodUs = 100000

	// What is the maximum cgroup depth we support?
	// We only expect to see a depth of around 3-4 at max, but we
	// allow 10 to give us some headroom. If this limit is reached
	// in a real world setting, we should increase it and ensure
	// valid headroom is given.
	maxSupportedCgroupDepth = 10
)

var errExceededMaxSupportedCgroupDepth = xerrors.Errorf("exceeded max supported cgroup depth of %d", maxSupportedCgroupDepth)

type cgroupV2Statter struct {
	parent *cgroupV2Statter
	path   string
	fs     afero.Fs
}

func newCgroupV2Statter(fs afero.Fs, path string, depth int) (*cgroupV2Statter, error) {
	var parent *cgroupV2Statter

	if depth >= maxSupportedCgroupDepth {
		return nil, errExceededMaxSupportedCgroupDepth
	}

	path = filepath.Clean(path)
	if parentPath := filepath.Dir(path); parentPath != path {
		var err error
		parent, err = newCgroupV2Statter(fs, parentPath, depth+1)
		if err != nil {
			return nil, err
		}
	}

	return &cgroupV2Statter{
		parent: parent,
		path:   filepath.Join(cgroupRootPath, path),
		fs:     fs,
	}, nil
}

// getFromParentOrFallback attempts to get a value from the parent cgroup
// when the current cgroup has no value set. If there is no parent, it
// calls the fallback function to determine the return value.
func getFromParentOrFallback[T any](
	s *cgroupV2Statter,
	getter func(*cgroupV2Statter) (T, error),
	fallback func() (T, error),
) (T, error) {
	if s.parent != nil {
		result, err := getter(s.parent)
		if err != nil {
			var zero T
			return zero, xerrors.Errorf("read parent: %w", err)
		}
		return result, nil
	}
	return fallback()
}

// cpuCounters returns one counter for every cgroup, from s up to the
// root, whose cpu.max sets a quota. The kernel enforces each of these
// quotas, so the caller reports whichever has the least room left.
//
// Each counter reads usage from the cgroup that sets the quota. cpu.stat
// includes all descendants, so this counts sibling cgroups that share the
// quota (for example /docker/<id> next to /init.scope in docker-in-docker).
//
// If no cgroup sets a quota, it returns a single unlimited counter for s.
func (s cgroupV2Statter) cpuCounters() ([]cpuCounter, error) {
	var counters []cpuCounter
	for c := &s; c != nil; c = c.parent {
		quotaUs, periodUs, ok, err := c.cpuMaxLocal()
		if err != nil {
			return nil, err
		}
		if !ok {
			continue
		}

		counters = append(counters, cpuCounter{
			total: quotaUs / periodUs,
			used: func() (float64, error) {
				usageUs, err := s.cpuUsageUsAt(c)
				if err != nil {
					return 0, err
				}
				return float64(usageUs) / periodUs, nil
			},
		})
	}
	if len(counters) > 0 {
		return counters, nil
	}

	periodUs, err := s.cpuPeriod()
	if err != nil {
		return nil, xerrors.Errorf("get cpu period: %w", err)
	}

	return []cpuCounter{{
		used: func() (float64, error) {
			usageUs, err := s.cpuUsageUsLocal()
			if err != nil {
				return 0, err
			}
			return float64(usageUs) / periodUs, nil
		},
	}}, nil
}

// cpuMaxLocal reads this cgroup's own cpu.max. ok is false when the file
// does not exist or the quota is "max".
func (s cgroupV2Statter) cpuMaxLocal() (quotaUs, periodUs float64, ok bool, err error) {
	cpuMaxPath := filepath.Join(s.path, cgroupV2CPUMax)

	quota, err := readInt64SepIdx(s.fs, cpuMaxPath, " ", 0)
	if err != nil {
		if !errors.Is(err, strconv.ErrSyntax) && !errors.Is(err, fs.ErrNotExist) {
			return 0, 0, false, xerrors.Errorf("get cpu quota: %w", err)
		}
		return 0, 0, false, nil
	}

	period, err := readInt64SepIdx(s.fs, cpuMaxPath, " ", 1)
	if err != nil {
		return 0, 0, false, xerrors.Errorf("get cpu period: %w", err)
	}

	return float64(quota), float64(period), true, nil
}

// cpuUsageUsAt returns usage_usec for the quota set on c. If c does not
// expose cpu.stat, it falls back to the process's own cgroup.
func (s cgroupV2Statter) cpuUsageUsAt(c *cgroupV2Statter) (int64, error) {
	if usageUs, err := c.cpuUsageUsLocal(); err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			return 0, xerrors.Errorf("read limit owner cpu: %w", err)
		}
		// The limit owner does not expose cpu.stat, so fall back to the
		// process's own cgroup below.
	} else {
		return usageUs, nil
	}

	return s.cpuUsageUsLocal()
}

func (s cgroupV2Statter) cpuUsageUsLocal() (int64, error) {
	cpuStatPath := filepath.Join(s.path, cgroupV2CPUStat)

	usageUs, err := readInt64Prefix(s.fs, cpuStatPath, "usage_usec")
	if err != nil {
		return 0, xerrors.Errorf("get cgroupv2 cpu used: %w", err)
	}

	return usageUs, nil
}

func (s cgroupV2Statter) cpuPeriod() (float64, error) {
	cpuMaxPath := filepath.Join(s.path, cgroupV2CPUMax)

	periodUs, err := readInt64SepIdx(s.fs, cpuMaxPath, " ", 1)
	if err != nil {
		if !errors.Is(err, strconv.ErrSyntax) && !errors.Is(err, fs.ErrNotExist) {
			return 0, xerrors.Errorf("get cpu period: %w", err)
		}

		// If the value is not a valid integer or the cpu.max file does
		// not exist, try the parent or return the default period. This can
		// happen in system-level cgroups like init.scope where cpu.max may not exist.
		return getFromParentOrFallback(&s, (*cgroupV2Statter).cpuPeriod, func() (float64, error) {
			return cgroupV2DefaultPeriodUs, nil
		})
	}

	return float64(periodUs), nil
}

// memoryMaxLocal reads this cgroup's own memory.max. ok is false when the
// file does not exist or the value is "max".
func (s cgroupV2Statter) memoryMaxLocal() (maxBytes float64, ok bool, err error) {
	memoryMaxPath := filepath.Join(s.path, cgroupV2MemoryMaxBytes)

	maxUsageBytes, err := readInt64(s.fs, memoryMaxPath)
	if err != nil {
		if !errors.Is(err, strconv.ErrSyntax) && !errors.Is(err, fs.ErrNotExist) {
			return 0, false, xerrors.Errorf("read memory total: %w", err)
		}
		return 0, false, nil
	}

	return float64(maxUsageBytes), true, nil
}

// memoryUsedAt returns memory.current minus inactive_file for the limit
// set on c. memory.current includes all descendants, so this counts
// sibling cgroups that share the limit. If c does not expose these files,
// it falls back to the process's own cgroup.
func (s cgroupV2Statter) memoryUsedAt(c *cgroupV2Statter) (float64, error) {
	if used, err := c.memoryUsedLocal(); err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			return 0, xerrors.Errorf("read limit owner memory: %w", err)
		}
		// The limit owner does not expose usage files, so fall back to
		// the process's own cgroup below.
	} else {
		return float64(used), nil
	}

	used, err := s.memoryUsed()
	if err != nil {
		return 0, err
	}
	return float64(used), nil
}

// memoryUsedLocal reads usage from this cgroup's own files only.
func (s cgroupV2Statter) memoryUsedLocal() (int64, error) {
	currUsageBytes, err := readInt64(s.fs, filepath.Join(s.path, cgroupV2MemoryUsageBytes))
	if err != nil {
		return 0, xerrors.Errorf("read memory current: %w", err)
	}

	inactiveFileBytes, err := readInt64Prefix(s.fs, filepath.Join(s.path, cgroupV2MemoryStat), "inactive_file")
	if err != nil {
		return 0, xerrors.Errorf("read memory stat inactive_file: %w", err)
	}

	return currUsageBytes - inactiveFileBytes, nil
}

// memoryUsed reads usage for this cgroup, taking each file from the
// nearest cgroup that has it.
func (s cgroupV2Statter) memoryUsed() (int64, error) {
	currUsageBytes, err := s.memoryCurrentBytes()
	if err != nil {
		return 0, xerrors.Errorf("read memory usage: %w", err)
	}

	inactiveFileBytes, err := s.memoryInactiveFileBytes()
	if err != nil {
		return 0, xerrors.Errorf("read memory stats: %w", err)
	}

	return currUsageBytes - inactiveFileBytes, nil
}

func (s cgroupV2Statter) memoryCurrentBytes() (int64, error) {
	memoryUsagePath := filepath.Join(s.path, cgroupV2MemoryUsageBytes)

	currUsageBytes, err := readInt64(s.fs, memoryUsagePath)
	if err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			return 0, xerrors.Errorf("read memory current: %w", err)
		}

		// If the memory current file does not exist, try the parent or return error.
		return getFromParentOrFallback(&s, (*cgroupV2Statter).memoryCurrentBytes, func() (int64, error) {
			return 0, xerrors.Errorf("read memory current: %w", err)
		})
	}

	return currUsageBytes, nil
}

func (s cgroupV2Statter) memoryInactiveFileBytes() (int64, error) {
	memoryStatPath := filepath.Join(s.path, cgroupV2MemoryStat)

	inactiveFileBytes, err := readInt64Prefix(s.fs, memoryStatPath, "inactive_file")
	if err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			return 0, xerrors.Errorf("read memory stat inactive_file: %w", err)
		}

		// If the memory stat file does not exist, try the parent or return error.
		return getFromParentOrFallback(&s, (*cgroupV2Statter).memoryInactiveFileBytes, func() (int64, error) {
			return 0, xerrors.Errorf("read memory stat inactive_file: %w", err)
		})
	}

	return inactiveFileBytes, nil
}

func (s cgroupV2Statter) memory(p Prefix) (*Result, error) {
	// https://docs.kernel.org/6.17/admin-guide/cgroup-v2.html#memory-interface-files
	r := &Result{
		Unit:   "B",
		Prefix: p,
	}

	// The kernel charges memory to every cgroup from s up to the root and
	// fails the allocation if any of them is over its memory.max. Report
	// the limit with the least room left, since that one is hit first.
	// Ties go to the cgroup nearest to s.
	var headroom float64
	for c := &s; c != nil; c = c.parent {
		limit, ok, err := c.memoryMaxLocal()
		if err != nil {
			return nil, err
		}
		if !ok {
			continue
		}

		used, err := s.memoryUsedAt(c)
		if err != nil {
			return nil, err
		}

		if r.Total == nil || limit-used < headroom {
			r.Total = ptr.To(limit)
			r.Used = used
			headroom = limit - used
		}
	}
	if r.Total != nil {
		return r, nil
	}

	used, err := s.memoryUsed()
	if err != nil {
		return nil, err
	}

	r.Used = float64(used)
	return r, nil
}
