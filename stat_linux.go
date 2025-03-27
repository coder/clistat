package clistat

import (
	"runtime"
	"strings"

	"github.com/spf13/afero"
)

func numCPU(fs afero.Fs) int {
	procCPUInfo, err := afero.ReadFile(fs, "/proc/cpuinfo")
	if err != nil {
		// If we have an issue reading from `/proc/cpuinfo`
		// we instead fallback to `runtime.NumCPU`.
		return runtime.NumCPU()
	}

	cpuInfo := strings.Split(string(procCPUInfo), "\n")
	cpuCount := 0

	for _, line := range cpuInfo {
		if line == "" {
			continue
		}

		parts := strings.Split(line, ":")

		if len(parts) > 1 && strings.TrimSpace(parts[0]) == "processor" {
			cpuCount += 1
		}
	}

	return cpuCount
}
