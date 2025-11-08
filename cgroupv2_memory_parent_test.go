package clistat

import (
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
)

// TestMemoryCurrentParentLookup verifies that memory.current falls back to parent when file doesn't exist
func TestMemoryCurrentParentLookup(t *testing.T) {
	t.Parallel()

	fs := afero.NewMemMapFs()
	
	// Setup: child cgroup path without memory.current, parent has it
	childPath := "/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod123.slice"
	parentPath := "/kubepods.slice/kubepods-burstable.slice"
	
	// Create parent cgroup with memory.current
	require.NoError(t, fs.MkdirAll(filepath.Join(cgroupRootPath, parentPath), 0755))
	require.NoError(t, afero.WriteFile(fs, 
		filepath.Join(cgroupRootPath, parentPath, cgroupV2MemoryUsageBytes), 
		[]byte("536870912"), 0644))
	require.NoError(t, afero.WriteFile(fs, 
		filepath.Join(cgroupRootPath, parentPath, cgroupV2MemoryStat), 
		[]byte("inactive_file 268435456"), 0644))
	require.NoError(t, afero.WriteFile(fs, 
		filepath.Join(cgroupRootPath, parentPath, cgroupV2MemoryMaxBytes), 
		[]byte("max"), 0644))
	
	// Create child directory but NO memory.current or memory.stat files
	require.NoError(t, fs.MkdirAll(filepath.Join(cgroupRootPath, childPath), 0755))
	
	// Create statter for child (which should recursively create parent)
	statter, err := newCgroupV2Statter(fs, childPath, 1)
	require.NoError(t, err)
	require.NotNil(t, statter)
	
	// This should succeed by looking up parent
	result, err := statter.memory(PrefixDefault)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, float64(536870912-268435456), result.Used)
}

// TestMemoryStatParentLookup verifies that memory.stat falls back to parent when file doesn't exist
func TestMemoryStatParentLookup(t *testing.T) {
	t.Parallel()

	fs := afero.NewMemMapFs()
	
	// Setup: child cgroup path without memory.stat, parent has it
	childPath := "/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod123.slice"
	parentPath := "/kubepods.slice/kubepods-burstable.slice"
	
	// Create parent cgroup with memory.stat but no memory.current
	require.NoError(t, fs.MkdirAll(filepath.Join(cgroupRootPath, parentPath), 0755))
	require.NoError(t, afero.WriteFile(fs, 
		filepath.Join(cgroupRootPath, parentPath, cgroupV2MemoryStat), 
		[]byte("inactive_file 268435456"), 0644))
	require.NoError(t, afero.WriteFile(fs, 
		filepath.Join(cgroupRootPath, parentPath, cgroupV2MemoryUsageBytes), 
		[]byte("536870912"), 0644))
	require.NoError(t, afero.WriteFile(fs, 
		filepath.Join(cgroupRootPath, parentPath, cgroupV2MemoryMaxBytes), 
		[]byte("max"), 0644))
	
	// Create child directory with memory.current but NO memory.stat
	require.NoError(t, fs.MkdirAll(filepath.Join(cgroupRootPath, childPath), 0755))
	require.NoError(t, afero.WriteFile(fs, 
		filepath.Join(cgroupRootPath, childPath, cgroupV2MemoryUsageBytes), 
		[]byte("536870912"), 0644))
	
	// Create statter for child
	statter, err := newCgroupV2Statter(fs, childPath, 1)
	require.NoError(t, err)
	require.NotNil(t, statter)
	
	// This should succeed by looking up memory.stat from parent
	result, err := statter.memory(PrefixDefault)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, float64(536870912-268435456), result.Used)
}
