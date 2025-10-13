package clistat

import (
	"syscall"
)

func isCGroupV2(path string) bool {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return false
	}
	return stat.Type == cgroupV2MagicNumber
}
