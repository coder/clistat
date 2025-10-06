package clistat

import (
	"errors"
	"io/fs"
	"strconv"

	"github.com/spf13/afero"
	"golang.org/x/xerrors"
	"tailscale.com/types/ptr"
)

// Paths for CGroupV2.
// Ref: https://docs.kernel.org/admin-guide/cgroup-v2.html
const (
	// Contains quota and period in microseconds separated by a space.
	cgroupV2CPUMax = "/sys/fs/cgroup/cpu.max"
	// Contains current CPU usage under usage_usec
	cgroupV2CPUStat = "/sys/fs/cgroup/cpu.stat"
	// Contains current cgroup memory usage in bytes.
	cgroupV2MemoryUsageBytes = "/sys/fs/cgroup/memory.current"
	// Contains max cgroup memory usage in bytes.
	cgroupV2MemoryMaxBytes = "/sys/fs/cgroup/memory.max"
	// Other memory stats - we are interested in total_inactive_file
	cgroupV2MemoryStat = "/sys/fs/cgroup/memory.stat"

	// https://www.kernel.org/doc/html/latest/admin-guide/cgroup-v2.html#cpu
	cgroupV2DefaultCPUPeriod = 100_000
)

type cgroupV2Statter struct {
	fs afero.Fs
}

func (s cgroupV2Statter) cpuUsed() (used float64, err error) {
	usageUs, err := readInt64Prefix(s.fs, cgroupV2CPUStat, "usage_usec")
	if err != nil {
		return 0, xerrors.Errorf("get cgroupv2 cpu used: %w", err)
	}
	periodUs, err := s.cpuPeriod()
	if err != nil {
		return 0, xerrors.Errorf("get cpu period: %w", err)
	}

	return float64(usageUs) / periodUs, nil
}

func (s cgroupV2Statter) cpuTotal() (total float64, err error) {
	periodUs, err := s.cpuPeriod()
	if err != nil {
		return 0, xerrors.Errorf("get cpu period: %w", err)
	}

	quotaUs, err := readInt64SepIdx(s.fs, cgroupV2CPUMax, " ", 0)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) || errors.Is(err, strconv.ErrSyntax) {
			// If the value is not a valid integer, assume it is the string
			// 'max' and that there is no limit set.
			return -1, nil
		}
		return 0, xerrors.Errorf("get cpu quota: %w", err)
	}

	return float64(quotaUs) / float64(periodUs), nil
}

func (s cgroupV2Statter) cpuPeriod() (float64, error) {
	// https://www.kernel.org/doc/html/latest/admin-guide/cgroup-v2.html#cpu
	periodUs, err := readInt64SepIdx(s.fs, cgroupV2CPUMax, " ", 1)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return float64(cgroupV2DefaultCPUPeriod), nil
		}
		return 0, xerrors.Errorf("get cpu period: %w", err)
	}
	return float64(periodUs), nil
}

func (s cgroupV2Statter) memory(p Prefix) (*Result, error) {
	r := &Result{
		Unit:   "B",
		Prefix: p,
	}
	maxUsageBytes, err := readInt64(s.fs, cgroupV2MemoryMaxBytes)
	if err != nil {
		if !errors.Is(err, strconv.ErrSyntax) {
			return nil, xerrors.Errorf("read memory total: %w", err)
		}
		// If the value is not a valid integer, assume it is the string
		// 'max' and that there is no limit set.
	} else {
		r.Total = ptr.To(float64(maxUsageBytes))
	}

	currUsageBytes, err := readInt64(s.fs, cgroupV2MemoryUsageBytes)
	if err != nil {
		return nil, xerrors.Errorf("read memory usage: %w", err)
	}

	inactiveFileBytes, err := readInt64Prefix(s.fs, cgroupV2MemoryStat, "inactive_file")
	if err != nil {
		return nil, xerrors.Errorf("read memory stats: %w", err)
	}

	r.Used = float64(currUsageBytes - inactiveFileBytes)
	return r, nil
}
