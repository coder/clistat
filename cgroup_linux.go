package clistat

import (
	"syscall"

	"github.com/spf13/afero"
)

func (s *Statter) isCGroupV2() bool {
	// If the underlying file system is an `OsFs`, then we will
	// make a `statfs` syscall to figure out if the filesystem
	// is cgroup v2 or not. This is unfortunately required as
	// afero doesn't implement this syscall functionality for us.
	if _, ok := s.fs.(*afero.OsFs); ok {
		var stat syscall.Statfs_t
		if err := syscall.Statfs(cgroupPath, &stat); err != nil {
			return false
		}

		return stat.Type == cgroupV2MagicNumber
	}

	// As a fall back, we will for the presence of /sys/fs/cgroup/cpu.max
	// NOTE(DanielleMaywood):
	// There is no requirement that a cgroup v2 file system will contain
	// this file, meaning this isn't completely foolproof.
	_, err := s.fs.Stat(cgroupV2CPUMax)
	return err == nil
}
