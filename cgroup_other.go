//go:build !linux

package clistat

func (s *Statter) isCGroupV2() bool {
	return false
}
