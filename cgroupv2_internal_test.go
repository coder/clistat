package clistat

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
)

func TestRecursiveCreation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		path        string
		startDepth  int
		expectError bool
	}{
		{
			path:        "",
			expectError: false,
		},
		{
			path:        "/",
			expectError: false,
		},
		{
			path:        "/a/small/path",
			expectError: false,
		},
		{
			path:        "/a/very/large/depth/that/should/absolutely/fail/because/it/is/too/deep",
			expectError: true,
		},
		{
			startDepth:  maxSupportCgroupDepth + 1,
			path:        "/a/very/large/depth/that/should/absolutely/fail/because/it/is/too/deep",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			t.Parallel()

			fs := afero.NewMemMapFs()
			statter, err := newCgroupV2Statter(fs, tt.path, tt.startDepth)
			if tt.expectError {
				require.ErrorIs(t, err, errExceededMaxSupportedCgroupDepth)
				require.Nil(t, statter)
			} else {
				require.NoError(t, err)
				require.NotNil(t, statter)
			}
		})
	}
}
