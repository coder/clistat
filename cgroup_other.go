//go:build !linux

package clistat

func isCgroupV2(path string) bool {
	return false
}
