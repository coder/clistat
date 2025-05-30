package clistat

import (
	"bufio"
	"bytes"
	"context"
	"strconv"
	"strings"

	"github.com/spf13/afero"
	"golang.org/x/xerrors"
	"tailscale.com/types/ptr"
)

// Path for cgroup
const (
	cgroupPath = "/sys/fs/cgroup"
)

const (
	// 0x63677270 (ascii for 'cgrp') is the magic number for identifying a cgroup v2
	// filesystem.
	// Ref: https://docs.kernel.org/admin-guide/cgroup-v2.html#mounting
	cgroupV2MagicNumber = 0x63677270
)

type cgroupStatter interface {
	cpuUsed() (used float64, err error)
	cpuTotal() (total float64, err error)
	memory(p Prefix) (*Result, error)
}

func (s *Statter) getCGroupStatter() cgroupStatter {
	if ok, err := s.IsContainerized(); err != nil || !ok {
		return nil
	}

	if s.isCGroupV2() {
		return &cgroupV2Statter{fs: s.fs}
	}

	return &cgroupV1Statter{fs: s.fs}
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
	// If the underlying file system is an `OsFs`, then we will
	// make a `statfs` syscall to figure out if the filesystem
	// is cgroup v2 or not. This is unfortunately required as
	// afero doesn't implement this syscall functionality for us.
	if _, ok := s.fs.(*afero.OsFs); ok {
		return isCGroupV2(cgroupPath)
	}

	s.logger.Debug(context.Background(), "not an *afero.OsFs, falling back to file existence check")

	// As a fall back, we will check for the presence of /sys/fs/cgroup/cpu.max
	// NOTE(DanielleMaywood):
	// There is no requirement that a cgroup v2 file system will contain
	// this file, meaning this isn't completely foolproof.
	_, err := s.fs.Stat(cgroupV2CPUMax)
	return err == nil
}

// ContainerMemory returns the memory usage of the container cgroup.
// If the system is not containerized, this always returns nil.
func (s *Statter) ContainerMemory(p Prefix) (*Result, error) {
	if s.cgroupStatter == nil {
		return nil, nil //nolint: nilnil
	}

	return s.cgroupStatter.memory(p)
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
