package clistat

import (
	"errors"
	"strconv"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/spf13/afero"
	"golang.org/x/xerrors"
	"tailscale.com/types/ptr"
)

// Paths for CGroupV1.
// Ref: https://www.kernel.org/doc/Documentation/cgroup-v1/cpuacct.txt
const (
	// CPU usage of all tasks in cgroup in nanoseconds.
	cgroupV1CPUAcctUsage = "/sys/fs/cgroup/cpu,cpuacct/cpuacct.usage"
	// CFS quota and period for cgroup in MICROseconds
	cgroupV1CFSQuotaUs = "/sys/fs/cgroup/cpu,cpuacct/cpu.cfs_quota_us"
	// CFS period for cgroup in MICROseconds
	cgroupV1CFSPeriodUs = "/sys/fs/cgroup/cpu,cpuacct/cpu.cfs_period_us"
	// Maximum memory usable by cgroup in bytes
	cgroupV1MemoryMaxUsageBytes = "/sys/fs/cgroup/memory/memory.limit_in_bytes"
	// Current memory usage of cgroup in bytes
	cgroupV1MemoryUsageBytes = "/sys/fs/cgroup/memory/memory.usage_in_bytes"
	// Other memory stats - we are interested in total_inactive_file
	cgroupV1MemoryStat = "/sys/fs/cgroup/memory/memory.stat"
)

const (
	// 9223372036854771712 is the highest positive signed 64-bit integer (263-1),
	// rounded down to multiples of 4096 (2^12), the most common page size on x86 systems.
	// This is used by docker to indicate no memory limit.
	UnlimitedMemory int64 = 9223372036854771712
)

type cgroupV1Statter struct {
	fs afero.Fs
}

func (s cgroupV1Statter) cpuUsed() (used float64, err error) {
	usageNs, err := readInt64(s.fs, cgroupV1CPUAcctUsage)
	if err != nil {
		// Try alternate path under /sys/fs/cgroup/cpuacct
		var merr error
		merr = multierror.Append(merr, xerrors.Errorf("read cpu used: %w", err))
		usageNs, err = readInt64(s.fs, strings.Replace(cgroupV1CPUAcctUsage, "cpu,cpuacct", "cpuacct", 1))
		if err != nil {
			merr = multierror.Append(merr, xerrors.Errorf("read cpu used: %w", err))
			return 0, merr
		}
	}

	// usage is in ns, convert to us
	usageNs /= 1000
	periodUs, err := readInt64(s.fs, cgroupV1CFSPeriodUs)
	if err != nil {
		// Try alternate path under /sys/fs/cpu
		var merr error
		merr = multierror.Append(merr, xerrors.Errorf("get cpu period: %w", err))
		periodUs, err = readInt64(s.fs, strings.Replace(cgroupV1CFSPeriodUs, "cpu,cpuacct", "cpu", 1))
		if err != nil {
			merr = multierror.Append(merr, xerrors.Errorf("get cpu period: %w", err))
			return 0, merr
		}
	}

	return float64(usageNs) / float64(periodUs), nil
}

func (s cgroupV1Statter) cpuTotal() (total float64, err error) {
	periodUs, err := readInt64(s.fs, cgroupV1CFSPeriodUs)
	if err != nil {
		// Try alternate path under /sys/fs/cpu
		var merr error
		merr = multierror.Append(merr, xerrors.Errorf("get cpu period: %w", err))
		periodUs, err = readInt64(s.fs, strings.Replace(cgroupV1CFSPeriodUs, "cpu,cpuacct", "cpu", 1))
		if err != nil {
			merr = multierror.Append(merr, xerrors.Errorf("get cpu period: %w", err))
			return 0, merr
		}
	}

	quotaUs, err := readInt64(s.fs, cgroupV1CFSQuotaUs)
	if err != nil {
		// Try alternate path under /sys/fs/cpu
		var merr error
		merr = multierror.Append(merr, xerrors.Errorf("get cpu quota: %w", err))
		quotaUs, err = readInt64(s.fs, strings.Replace(cgroupV1CFSQuotaUs, "cpu,cpuacct", "cpu", 1))
		if err != nil {
			merr = multierror.Append(merr, xerrors.Errorf("get cpu quota: %w", err))
			return 0, merr
		}
	}

	if quotaUs < 0 {
		return -1, nil
	}

	return float64(quotaUs) / float64(periodUs), nil
}

func (s cgroupV1Statter) memory(p Prefix) (*Result, error) {
	r := &Result{
		Unit:   "B",
		Prefix: p,
	}
	maxUsageBytes, err := readInt64(s.fs, cgroupV1MemoryMaxUsageBytes)
	if err != nil {
		if !errors.Is(err, strconv.ErrSyntax) {
			return nil, xerrors.Errorf("read memory total: %w", err)
		}
		// I haven't found an instance where this isn't a valid integer.
		// Nonetheless, if it is not, assume there is no limit set.
		maxUsageBytes = -1
	}
	// Set to unlimited if we detect the unlimited docker value.
	if maxUsageBytes == UnlimitedMemory {
		maxUsageBytes = -1
	}

	// need a space after total_rss so we don't hit something else
	usageBytes, err := readInt64(s.fs, cgroupV1MemoryUsageBytes)
	if err != nil {
		return nil, xerrors.Errorf("read memory usage: %w", err)
	}

	totalInactiveFileBytes, err := readInt64Prefix(s.fs, cgroupV1MemoryStat, "total_inactive_file")
	if err != nil {
		return nil, xerrors.Errorf("read memory stats: %w", err)
	}

	// If max usage bytes is -1, there is no memory limit set.
	if maxUsageBytes > 0 {
		r.Total = ptr.To(float64(maxUsageBytes))
	}

	// Total memory used is usage - total_inactive_file
	r.Used = float64(usageBytes - totalInactiveFileBytes)

	return r, nil
}
