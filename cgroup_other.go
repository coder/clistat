//go:build !linux

package clistat

func isCgroupV2(_ string) bool {
	return false
}
