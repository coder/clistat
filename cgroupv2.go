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

func (s cgroupV2Statter) cpuUsed() (used float64, err error) {
	cpuStatPath := filepath.Join(s.path, cgroupV2CPUStat)

	usageUs, err := readInt64Prefix(s.fs, cpuStatPath, "usage_usec")
	if err != nil {
		return 0, xerrors.Errorf("get cgroupv2 cpu used: %w", err)
	}
	periodUs, err := s.cpuPeriod()
	if err != nil {
		return 0, xerrors.Errorf("get cpu period: %w", err)
	}

	return float64(usageUs) / periodUs, nil
}

func (s cgroupV2Statter) cpuQuota() (float64, error) {
	cpuMaxPath := filepath.Join(s.path, cgroupV2CPUMax)

	quotaUs, err := readInt64SepIdx(s.fs, cpuMaxPath, " ", 0)
	if err != nil {
		if !errors.Is(err, strconv.ErrSyntax) && !errors.Is(err, fs.ErrNotExist) {
			return 0, xerrors.Errorf("get cpu quota: %w", err)
		}

		// If the value is not a valid integer, assume it is the string
		// 'max' and that there is no limit set. Try the parent or return -1.
		return getFromParentOrFallback(&s, (*cgroupV2Statter).cpuQuota, func() (float64, error) {
			return -1.0, nil
		})
	}

	return float64(quotaUs), nil
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

func (s cgroupV2Statter) cpuTotal() (total float64, err error) {
	periodUs, err := s.cpuPeriod()
	if err != nil {
		return 0, xerrors.Errorf("get cpu period: %w", err)
	}

	quotaUs, err := s.cpuQuota()
	if err != nil {
		return 0, xerrors.Errorf("get cpu quota: %w", err)
	}

	return quotaUs / periodUs, nil
}

// memoryLimit is the effective memory.max and the cgroup that sets it.
// owner is kept so usage can be read from the same cgroup the limit is
// enforced on.
type memoryLimit struct {
	owner *cgroupV2Statter
	bytes float64
}

// memoryLimit returns the nearest cgroup, starting at s and walking up,
// that sets memory.max. It returns nil when no cgroup sets a limit.
func (s *cgroupV2Statter) memoryLimit() (*memoryLimit, error) {
	memoryMaxPath := filepath.Join(s.path, cgroupV2MemoryMaxBytes)

	maxUsageBytes, err := readInt64(s.fs, memoryMaxPath)
	if err != nil {
		if !errors.Is(err, strconv.ErrSyntax) && !errors.Is(err, fs.ErrNotExist) {
			return nil, xerrors.Errorf("read memory total: %w", err)
		}

		// If the value is not a valid integer _or_ the memory max file
		// does not exist, we can assume that the limit is 'max'.
		// Try the parent or return nil (no limit).
		return getFromParentOrFallback(s, (*cgroupV2Statter).memoryLimit, func() (*memoryLimit, error) {
			return nil, nil
		})
	}

	return &memoryLimit{owner: s, bytes: float64(maxUsageBytes)}, nil
}

func (s cgroupV2Statter) memoryUsedBytes() (int64, error) {
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
	limit, err := s.memoryLimit()
	if err != nil {
		return nil, xerrors.Errorf("read memory total: %w", err)
	}

	// Read usage from the cgroup that owns the limit so that used and
	// total describe the same domain. Otherwise sibling cgroups under
	// the limit owner (for example /docker/<id> next to /init.scope in
	// docker-in-docker) would count against the limit but not the usage.
	if limit != nil {
		r.Total = ptr.To(limit.bytes)

		if limit.owner.path != s.path {
			if used, err := limit.owner.memoryUsedBytes(); err != nil {
				if !errors.Is(err, fs.ErrNotExist) {
					return nil, xerrors.Errorf("read limit owner memory: %w", err)
				}
				// The limit owner does not expose usage files, so fall
				// back to the process's own cgroup below.
			} else {
				r.Used = float64(used)
				return r, nil
			}
		}
	}

	used, err := s.memoryUsedBytes()
	if err != nil {
		return nil, err
	}

	r.Used = float64(used)
	return r, nil
}
