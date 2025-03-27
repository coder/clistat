//go:build !linux

package clistat

import (
	"runtime"
)

func (s *Statter) numCPU() int {
	return runtime.NumCPU()
}
