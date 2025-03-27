//go:build linux

package clistat

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLinuxNumCPU(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		cpuInfo       string
		expectedCount int
	}{
		{
			name:          "Raspberry Pi 5 (4 CPU)",
			cpuInfo:       procCPUInfoRaspberryPi5,
			expectedCount: 4,
		},
		{
			name:          "EU Coder Workspace (96 CPU)",
			cpuInfo:       procCPUInfoCoderWorkspaceEU,
			expectedCount: 96,
		},
		{
			name:          "UpCloud Cloud Server (1 CPU)",
			cpuInfo:       procCPUInfoUpCloud,
			expectedCount: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			fs := initFS(t, map[string]string{
				"/proc/cpuinfo": tt.cpuInfo,
			})

			s, err := New(WithFS(fs))
			require.NoError(t, err)

			cpuCount := s.numCPU()
			require.Equal(t, tt.expectedCount, cpuCount)
		})
	}
}

//go:embed testdata/cpuinfo/raspberrypi5
var procCPUInfoRaspberryPi5 string

//go:embed testdata/cpuinfo/coderworkspaceeu
var procCPUInfoCoderWorkspaceEU string

//go:embed testdata/cpuinfo/upcloud
var procCPUInfoUpCloud string
