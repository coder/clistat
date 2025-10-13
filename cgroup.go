package clistat

import (
	"bufio"
	"bytes"
	"errors"
	"strconv"
	"strings"

	"github.com/spf13/afero"
	"golang.org/x/xerrors"
	"tailscale.com/types/ptr"
)

// Paths for cgroup
const (
	cgroupRootPath = "/sys/fs/cgroup"

	// Contains a path to the cgroup
	procSelfCgroup = "/proc/self/cgroup"
)

const (
	// 0x63677270 (ascii for 'cgrp') is the magic number for identifying a cgroup v2
	// filesystem.
	// Ref: https://docs.kernel.org/admin-guide/cgroup-v2.html#mounting
	cgroupV2MagicNumber = 0x63677270
)

var errNotContainerized = errors.New("not containerized")

type cgroupStatter interface {
	cpuUsed() (used float64, err error)
	cpuTotal() (total float64, err error)
	memory(p Prefix) (*Result, error)
}

func (s *Statter) getCGroupStatter() (cgroupStatter, error) {
	isContainerized, err := s.IsContainerized()
	if err != nil {
		return nil, err
	}
	if !isContainerized {
		return nil, errNotContainerized
	}

	if s.isCGroupV2() {
		cgroupPath, err := currentProcCgroup(s.fs)
		if err != nil {
			return nil, xerrors.Errorf("get current cgroup: %w", err)
		}

		return newCgroupV2Statter(s.fs, cgroupPath, 1)
	}

	return &cgroupV1Statter{fs: s.fs}, nil
}

// ContainerCPU returns the CPU usage of the container cgroup.
// This is calculated as difference of two samples of the
// CPU usage of the container cgroup.
// The total is read from the relevant path in /sys/fs/cgroup.
// If there is no limit set, the total is assumed to be the
// number of host cores multiplied by the CFS period.
// If the system is not containerized, this always returns nil.
func (s *Statter) ContainerCPU() (*Result, error) {
	if s.cgroupStatter == nil {
		return nil, nil //nolint: nilnil
	}

	total, err := s.cgroupStatter.cpuTotal()
	if err != nil {
		return nil, xerrors.Errorf("get total cpu: %w", err)
	}
	used1, err := s.cgroupStatter.cpuUsed()
	if err != nil {
		return nil, xerrors.Errorf("get cgroup CPU usage: %w", err)
	}

	// The measurements in /sys/fs/cgroup are counters.
	// We need to wait for a bit to get a difference.
	// Note that someone could reset the counter in the meantime.
	// We can't do anything about that.
	s.wait(s.sampleInterval)

	used2, err := s.cgroupStatter.cpuUsed()
	if err != nil {
		return nil, xerrors.Errorf("get cgroup CPU usage: %w", err)
	}

	if used2 < used1 {
		// Someone reset the counter. Best we can do is count from zero.
		used1 = 0
	}

	r := &Result{
		Unit:   "cores",
		Used:   used2 - used1,
		Prefix: PrefixDefault,
	}

	if total > 0 {
		r.Total = ptr.To(total)
	}
	return r, nil
}

func (s *Statter) isCGroupV2() bool {
	return s.cgroupV2Detector(s.fs)
}

// ContainerMemory returns the memory usage of the container cgroup.
// If the system is not containerized, this always returns nil.
func (s *Statter) ContainerMemory(p Prefix) (*Result, error) {
	if s.cgroupStatter == nil {
		return nil, nil //nolint: nilnil
	}

	return s.cgroupStatter.memory(p)
}

func currentProcCgroup(fs afero.Fs) (string, error) {
	data, err := afero.ReadFile(fs, procSelfCgroup)
	if err != nil {
		return "", xerrors.Errorf("read %v: %w", procSelfCgroup, err)
	}

	parts := strings.Split(strings.TrimSpace(string(data)), ":")
	if len(parts) != 3 {
		return "", xerrors.Errorf("parse %v: %w", procSelfCgroup, err)
	}

	return parts[2], nil
}

// read an int64 value from path
func readInt64(fs afero.Fs, path string) (int64, error) {
	data, err := afero.ReadFile(fs, path)
	if err != nil {
		return 0, xerrors.Errorf("read %s: %w", path, err)
	}

	val, err := strconv.ParseInt(string(bytes.TrimSpace(data)), 10, 64)
	if err != nil {
		return 0, xerrors.Errorf("parse %s: %w", path, err)
	}

	return val, nil
}

// read an int64 value from path at field idx separated by sep
func readInt64SepIdx(fs afero.Fs, path, sep string, idx int) (int64, error) {
	data, err := afero.ReadFile(fs, path)
	if err != nil {
		return 0, xerrors.Errorf("read %s: %w", path, err)
	}

	parts := strings.Split(string(data), sep)
	if len(parts) < idx {
		return 0, xerrors.Errorf("expected line %q to have at least %d parts", string(data), idx+1)
	}

	val, err := strconv.ParseInt(strings.TrimSpace(parts[idx]), 10, 64)
	if err != nil {
		return 0, xerrors.Errorf("parse %s: %w", path, err)
	}

	return val, nil
}

// read the first int64 value from path prefixed with prefix
func readInt64Prefix(fs afero.Fs, path, prefix string) (int64, error) {
	data, err := afero.ReadFile(fs, path)
	if err != nil {
		return 0, xerrors.Errorf("read %s: %w", path, err)
	}

	scn := bufio.NewScanner(bytes.NewReader(data))
	for scn.Scan() {
		line := strings.TrimSpace(scn.Text())
		if !strings.HasPrefix(line, prefix) {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) != 2 {
			return 0, xerrors.Errorf("parse %s: expected two fields but got %s", path, line)
		}

		val, err := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)
		if err != nil {
			return 0, xerrors.Errorf("parse %s: %w", path, err)
		}

		return val, nil
	}

	return 0, xerrors.Errorf("parse %s: did not find line with prefix %s", path, prefix)
}
