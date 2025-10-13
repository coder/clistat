package clistat

import (
	"syscall"
)

func isCgroupV2(path string) bool {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return false
	}
	return stat.Type == cgroupV2MagicNumber
}
