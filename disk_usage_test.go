package clistat_test

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/coder/clistat"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDiskUsage(t *testing.T) {
	t.Parallel()

	s, err := clistat.New()
	require.NoError(t, err)

	t.Run("EmptyDirectory", func(t *testing.T) {
		t.Parallel()

		tmpDir := t.TempDir()

		result, err := s.DiskUsage(clistat.PrefixDefault, tmpDir)
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, float64(0), result.Used)
		assert.Nil(t, result.Total)
		assert.Equal(t, "B", result.Unit)
	})

	t.Run("DirectoryWithFiles", func(t *testing.T) {
		t.Parallel()

		tmpDir := t.TempDir()

		// Create some test files with known sizes
		file1 := filepath.Join(tmpDir, "file1.txt")
		file2 := filepath.Join(tmpDir, "file2.txt")
		subDir := filepath.Join(tmpDir, "subdir")
		file3 := filepath.Join(subDir, "file3.txt")

		require.NoError(t, os.WriteFile(file1, make([]byte, 1024), 0o644))
		require.NoError(t, os.WriteFile(file2, make([]byte, 2048), 0o644))
		require.NoError(t, os.MkdirAll(subDir, 0o755))
		require.NoError(t, os.WriteFile(file3, make([]byte, 4096), 0o644))

		result, err := s.DiskUsage(clistat.PrefixDefault, tmpDir)
		require.NoError(t, err)
		assert.NotNil(t, result)
		// The exact size might vary due to block allocation,
		// but it should be at least the sum of file sizes
		assert.GreaterOrEqual(t, result.Used, float64(1024+2048+4096))
		assert.Nil(t, result.Total)
		assert.Equal(t, "B", result.Unit)
	})

	t.Run("WithPrefix", func(t *testing.T) {
		t.Parallel()

		tmpDir := t.TempDir()
		file := filepath.Join(tmpDir, "file.txt")
		require.NoError(t, os.WriteFile(file, make([]byte, 1024), 0o644))

		result, err := s.DiskUsage(clistat.PrefixKibi, tmpDir)
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, clistat.Prefix(clistat.PrefixKibi), result.Prefix)
		// String representation should use KiB
		str := result.String()
		assert.Contains(t, str, "KiB")
	})

	t.Run("NonExistentPath", func(t *testing.T) {
		t.Parallel()

		result, err := s.DiskUsage(clistat.PrefixDefault, "/nonexistent/path/that/does/not/exist")
		// WalkDir may or may not return an error for non-existent root paths
		// depending on OS. Either an error or zero usage is acceptable.
		if err == nil {
			assert.NotNil(t, result)
			assert.Equal(t, float64(0), result.Used)
		}
	})

	t.Run("SymlinksAreSkipped", func(t *testing.T) {
		t.Parallel()

		if runtime.GOOS == "windows" {
			t.Skip("Symlink creation may require elevated privileges on Windows")
		}

		tmpDir := t.TempDir()

		// Create a real file
		realFile := filepath.Join(tmpDir, "real.txt")
		require.NoError(t, os.WriteFile(realFile, make([]byte, 1024), 0o644))

		// Create an external directory with a file
		externalDir := t.TempDir()
		externalFile := filepath.Join(externalDir, "external.txt")
		require.NoError(t, os.WriteFile(externalFile, make([]byte, 5000), 0o644))

		// Create a symlink to the external file (should be skipped)
		symlink := filepath.Join(tmpDir, "link_to_external.txt")
		require.NoError(t, os.Symlink(externalFile, symlink))

		// Create a symlink to the real file (should also be skipped)
		symlinkInternal := filepath.Join(tmpDir, "link_to_real.txt")
		require.NoError(t, os.Symlink(realFile, symlinkInternal))

		result, err := s.DiskUsage(clistat.PrefixDefault, tmpDir)
		require.NoError(t, err)
		assert.NotNil(t, result)

		// Should only count the real file, not the symlinks or their targets
		// The used space should be approximately 1024 bytes (the real file),
		// not 1024 + 5000 + 1024 (if symlinks were followed)
		// Using block size, so check it's less than what it would be with symlinks
		assert.Less(t, result.Used, float64(5000), "Symlinks should not be followed")
	})

	t.Run("SymlinkLoopsDoNotCauseInfiniteLoop", func(t *testing.T) {
		t.Parallel()

		if runtime.GOOS == "windows" {
			t.Skip("Symlink creation may require elevated privileges on Windows")
		}

		tmpDir := t.TempDir()

		// Create a file
		realFile := filepath.Join(tmpDir, "real.txt")
		require.NoError(t, os.WriteFile(realFile, make([]byte, 512), 0o644))

		// Create a symlink loop: dir/loop -> dir
		loopLink := filepath.Join(tmpDir, "loop")
		require.NoError(t, os.Symlink(tmpDir, loopLink))

		// This should complete without hanging
		result, err := s.DiskUsage(clistat.PrefixDefault, tmpDir)
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Greater(t, result.Used, float64(0))
	})

	t.Run("HardLinksCountedOnce", func(t *testing.T) {
		t.Parallel()

		if runtime.GOOS == "windows" {
			t.Skip("Hard link behavior differs on Windows")
		}

		tmpDir := t.TempDir()

		// Create a file
		originalFile := filepath.Join(tmpDir, "original.txt")
		require.NoError(t, os.WriteFile(originalFile, make([]byte, 4096), 0o644))

		// Create hard links to the same file
		hardLink1 := filepath.Join(tmpDir, "hardlink1.txt")
		hardLink2 := filepath.Join(tmpDir, "hardlink2.txt")
		require.NoError(t, os.Link(originalFile, hardLink1))
		require.NoError(t, os.Link(originalFile, hardLink2))

		result, err := s.DiskUsage(clistat.PrefixDefault, tmpDir)
		require.NoError(t, err)
		assert.NotNil(t, result)

		// With hard link deduplication, should count the file only once.
		// The result should be around 4096 bytes (one block), not 12288 (3x).
		// Account for block size variance
		assert.Less(t, result.Used, float64(8192), "Hard links should be deduplicated")
	})

	t.Run("PermissionDeniedSkipped", func(t *testing.T) {
		t.Parallel()

		if runtime.GOOS == "windows" {
			t.Skip("Permission handling differs on Windows")
		}

		if os.Getuid() == 0 {
			t.Skip("Test cannot run as root (root can read anything)")
		}

		tmpDir := t.TempDir()

		// Create an accessible file
		accessibleFile := filepath.Join(tmpDir, "accessible.txt")
		require.NoError(t, os.WriteFile(accessibleFile, make([]byte, 1024), 0o644))

		// Create an inaccessible subdirectory
		inaccessibleDir := filepath.Join(tmpDir, "noaccess")
		require.NoError(t, os.MkdirAll(inaccessibleDir, 0o755))
		inaccessibleFile := filepath.Join(inaccessibleDir, "secret.txt")
		require.NoError(t, os.WriteFile(inaccessibleFile, make([]byte, 5000), 0o644))

		// Remove read permission on the directory
		require.NoError(t, os.Chmod(inaccessibleDir, 0o000))
		t.Cleanup(func() {
			// Restore permissions for cleanup
			_ = os.Chmod(inaccessibleDir, 0o755)
		})

		// Should not error, just skip the inaccessible directory
		result, err := s.DiskUsage(clistat.PrefixDefault, tmpDir)
		require.NoError(t, err)
		assert.NotNil(t, result)

		// Should have counted the accessible file but not the inaccessible one
		assert.Greater(t, result.Used, float64(0))
		// If it counted everything, it would be at least 6000 bytes
		assert.Less(t, result.Used, float64(5000), "Inaccessible files should be skipped")
	})
}

func TestDiskUsageWithTotal(t *testing.T) {
	t.Parallel()

	s, err := clistat.New()
	require.NoError(t, err)

	t.Run("IncludesFilesystemTotal", func(t *testing.T) {
		t.Parallel()

		tmpDir := t.TempDir()
		file := filepath.Join(tmpDir, "file.txt")
		require.NoError(t, os.WriteFile(file, make([]byte, 1024), 0o644))

		result, err := s.DiskUsageWithTotal(clistat.PrefixDefault, tmpDir)
		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.GreaterOrEqual(t, result.Used, float64(1024))
		// Should have a Total from the filesystem
		assert.NotNil(t, result.Total)
		assert.Greater(t, *result.Total, float64(0))
		assert.Equal(t, "B", result.Unit)
	})
}

func TestDiskUsageSimple(t *testing.T) {
	t.Parallel()

	s, err := clistat.New()
	require.NoError(t, err)

	t.Run("UsesFileSizes", func(t *testing.T) {
		t.Parallel()

		tmpDir := t.TempDir()
		file := filepath.Join(tmpDir, "file.txt")
		// Write exactly 1000 bytes to test that Simple uses file size, not blocks
		require.NoError(t, os.WriteFile(file, make([]byte, 1000), 0o644))

		result, err := s.DiskUsageSimple(clistat.PrefixDefault, tmpDir)
		require.NoError(t, err)
		assert.NotNil(t, result)
		// Simple version should report exactly the file size
		assert.Equal(t, float64(1000), result.Used)
	})

	t.Run("SymlinksAreSkipped", func(t *testing.T) {
		t.Parallel()

		if runtime.GOOS == "windows" {
			t.Skip("Symlink creation may require elevated privileges on Windows")
		}

		tmpDir := t.TempDir()

		// Create a real file
		realFile := filepath.Join(tmpDir, "real.txt")
		require.NoError(t, os.WriteFile(realFile, make([]byte, 500), 0o644))

		// Create an external file
		externalDir := t.TempDir()
		externalFile := filepath.Join(externalDir, "external.txt")
		require.NoError(t, os.WriteFile(externalFile, make([]byte, 10000), 0o644))

		// Symlink to external file
		symlink := filepath.Join(tmpDir, "link.txt")
		require.NoError(t, os.Symlink(externalFile, symlink))

		result, err := s.DiskUsageSimple(clistat.PrefixDefault, tmpDir)
		require.NoError(t, err)
		assert.NotNil(t, result)

		// Should only count the real file (500 bytes)
		assert.Equal(t, float64(500), result.Used)
	})
}
