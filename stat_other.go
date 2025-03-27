//go:build !linux

package clistat

import (
	"runtime"

	"github.com/spf13/afero"
)

func numCPU(_ afero.Fs) int {
	return runtime.NumCPU()
}
