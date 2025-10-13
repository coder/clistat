//go:build !linux

package clistat

func isCGroupV2(path string) bool {
	return false
}
