package clistat

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCurrentProcCgroup(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		procFile    string
		expectError string
		expectPath  string
	}{
		{
			name:       "RootPath",
			procFile:   `0::/`,
			expectPath: "/",
		},
		{
			name:       "SingleLevelPath",
			procFile:   `0::/init.slice`,
			expectPath: "/init.slice",
		},
		{
			name:       "MultipleLevelSlice",
			procFile:   `0::/wibble/wobble/init.slice`,
			expectPath: "/wibble/wobble/init.slice",
		},
		{
			name: "MixOfHierachy",
			procFile: `1:net_cls:/
0::/`,
			expectPath: "/",
		},
		{
			name: "MixOfHierarchyPaths",
			procFile: `1:net_cls:/init.slice
0::/`,
			expectPath: "/",
		},
		{
			name: "MixOfHierarchyPaths/Order",
			procFile: `0::/
1:net_cls:/init.slice`,
			expectPath: "/",
		},
		{
			name: "MixOfHierarchyPaths/Paths",
			procFile: `0::/init.slice
1:net_cls:/`,
			expectPath: "/init.slice",
		},
		{
			name:        "MissingHierarchy0",
			procFile:    `1:net_cls:/`,
			expectError: "no cgroup entry for hierarchy 0 found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			fs := initFS(t, map[string]string{procSelfCgroup: tt.procFile})

			path, err := currentProcCgroup(fs)
			if tt.expectError != "" {
				require.ErrorContains(t, err, tt.expectError)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expectPath, path)
			}
		})
	}
}
