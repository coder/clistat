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

// Paths for CGroupV2.
// Ref: https://docs.kernel.org/admin-guide/cgroup-v2.html
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

func (s cgroupV2Statter) cpuUsed() (used float64, err error) {
	cpuStatPath := filepath.Join(s.path, cgroupV2CPUStat)
	cpuMaxPath := filepath.Join(s.path, cgroupV2CPUMax)

	usageUs, err := readInt64Prefix(s.fs, cpuStatPath, "usage_usec")
	if err != nil {
		return 0, xerrors.Errorf("get cgroupv2 cpu used: %w", err)
	}
	periodUs, err := readInt64SepIdx(s.fs, cpuMaxPath, " ", 1)
	if err != nil {
		return 0, xerrors.Errorf("get cpu period: %w", err)
	}

	return float64(usageUs) / float64(periodUs), nil
}

func (s cgroupV2Statter) cpuQuota() (float64, error) {
	cpuMaxPath := filepath.Join(s.path, cgroupV2CPUMax)

	quotaUs, err := readInt64SepIdx(s.fs, cpuMaxPath, " ", 0)
	if err != nil {
		if !errors.Is(err, strconv.ErrSyntax) && !errors.Is(err, fs.ErrNotExist) {
			return 0, xerrors.Errorf("get cpu quota: %w", err)
		}

		// If the value is not a valid integer, assume it is the string
		// 'max' and that there is no limit set. In this scenario, we call
		// the parent to find its quota.
		if s.parent != nil {
			total, err := s.parent.cpuQuota()
			if err != nil {
				return 0, xerrors.Errorf("get parent cpu quota: %w", err)
			}
			return total, nil
		}

		return -1, nil
	}

	return float64(quotaUs), nil
}

func (s cgroupV2Statter) cpuTotal() (total float64, err error) {
	cpuMaxPath := filepath.Join(s.path, cgroupV2CPUMax)

	periodUs, err := readInt64SepIdx(s.fs, cpuMaxPath, " ", 1)
	if err != nil {
		return 0, xerrors.Errorf("get cpu period: %w", err)
	}

	quotaUs, err := s.cpuQuota()
	if err != nil {
		return 0, xerrors.Errorf("get cpu quota: %w", err)
	}

	return float64(quotaUs) / float64(periodUs), nil
}

func (s cgroupV2Statter) memoryMaxBytes() (*float64, error) {
	memoryMaxPath := filepath.Join(s.path, cgroupV2MemoryMaxBytes)

	maxUsageBytes, err := readInt64(s.fs, memoryMaxPath)
	if err != nil {
		if !errors.Is(err, strconv.ErrSyntax) && !errors.Is(err, fs.ErrNotExist) {
			return nil, xerrors.Errorf("read memory total: %w", err)
		}

		// If the value is not a valid integer _or_ the memory max file
		// does not exist, than we can assume that the limit is 'max'.
		// If the memory limit is max, and we have a parent, we shall call
		// the parent to find its maximum memory value.
		if s.parent != nil {
			result, err := s.parent.memoryMaxBytes()
			if err != nil {
				return nil, xerrors.Errorf("read parent memory max: %w", err)
			}
			return result, nil
		}

		// We have no parent, and no max memory limit, so there is no memory limit.
		return nil, nil
	}

	return ptr.To(float64(maxUsageBytes)), nil
}

func (s cgroupV2Statter) memory(p Prefix) (*Result, error) {
	memoryUsagePath := filepath.Join(s.path, cgroupV2MemoryUsageBytes)
	memoryStatPath := filepath.Join(s.path, cgroupV2MemoryStat)

	// https://docs.kernel.org/admin-guide/cgroup-v2.html#memory-interface-files
	r := &Result{
		Unit:   "B",
		Prefix: p,
	}
	if total, err := s.memoryMaxBytes(); err != nil {
		return nil, xerrors.Errorf("read memory total: %w", err)
	} else {
		r.Total = total
	}

	currUsageBytes, err := readInt64(s.fs, memoryUsagePath)
	if err != nil {
		return nil, xerrors.Errorf("read memory usage: %w", err)
	}

	inactiveFileBytes, err := readInt64Prefix(s.fs, memoryStatPath, "inactive_file")
	if err != nil {
		return nil, xerrors.Errorf("read memory stats: %w", err)
	}

	r.Used = float64(currUsageBytes - inactiveFileBytes)
	return r, nil
}
