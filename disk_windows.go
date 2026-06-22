package clistat

import (
	"io/fs"
	"path/filepath"

	"golang.org/x/sys/windows"
	"tailscale.com/types/ptr"
)

// Disk returns the disk usage of the given path.
// If path is empty, it defaults to C:\
func (*Statter) Disk(p Prefix, path string) (*Result, error) {
	if path == "" {
		path = `C:\`
	}

	pathPtr, err := windows.UTF16PtrFromString(path)
	if err != nil {
		return nil, err
	}

	var freeBytes, totalBytes, availBytes uint64
	if err := windows.GetDiskFreeSpaceEx(
		pathPtr,
		&freeBytes,
		&totalBytes,
		&availBytes,
	); err != nil {
		return nil, err
	}

	var r Result
	r.Total = ptr.To(float64(totalBytes))
	r.Used = float64(totalBytes - freeBytes)
	r.Unit = "B"
	r.Prefix = p
	return &r, nil
}

// DiskUsage returns the actual disk usage of a directory tree,
// similar to "du -sh". This is useful in containerized environments
// where you want to track usage of specific directories rather than
// the entire filesystem.
//
// Unlike Disk(), which uses GetDiskFreeSpaceEx to get filesystem-level usage,
// DiskUsage walks the directory tree and sums up file sizes.
//
// Symlinks are not followed to avoid counting files outside the
// target directory and to prevent infinite loops from symlink cycles.
//
// Note: This operation can be expensive for large directory trees
// with many small files. Consider using appropriate refresh intervals.
// Files that cannot be accessed (permission errors, etc.) are skipped
// silently.
func (*Statter) DiskUsage(p Prefix, path string) (*Result, error) {
	if path == "" {
		path = `C:\`
	}

	var totalSize int64

	err := filepath.WalkDir(path, func(_ string, d fs.DirEntry, err error) error {
		if err != nil {
			// Skip files/directories we can't access
			return nil
		}

		// Skip directories themselves, we only count file sizes
		if d.IsDir() {
			return nil
		}

		// Skip symlinks to avoid:
		// 1. Counting files outside the target directory
		// 2. Infinite loops from symlink cycles
		// 3. Double-counting if symlink target is also in the tree
		if d.Type()&fs.ModeSymlink != 0 {
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return nil
		}

		totalSize += info.Size()
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &Result{
		Used:   float64(totalSize),
		Total:  nil, // Directory usage doesn't have a "total" concept
		Unit:   "B",
		Prefix: p,
	}, nil
}

// DiskUsageWithTotal returns the actual disk usage of a directory tree
// along with the total filesystem capacity. This combines DiskUsage
// with filesystem-level total from Disk.
func (s *Statter) DiskUsageWithTotal(p Prefix, path string) (*Result, error) {
	if path == "" {
		path = `C:\`
	}

	usage, err := s.DiskUsage(p, path)
	if err != nil {
		return nil, err
	}

	// Get the filesystem total for the path
	pathPtr, err := windows.UTF16PtrFromString(path)
	if err != nil {
		return usage, nil
	}

	var freeBytes, totalBytes, availBytes uint64
	if err := windows.GetDiskFreeSpaceEx(
		pathPtr,
		&freeBytes,
		&totalBytes,
		&availBytes,
	); err != nil {
		// Return usage without total if we can't get fs stats
		return usage, nil
	}

	usage.Total = ptr.To(float64(totalBytes))
	return usage, nil
}

// DiskUsageSimple is identical to DiskUsage on Windows.
// On Unix systems, DiskUsage uses disk blocks for accuracy
// while DiskUsageSimple uses file sizes. On Windows, both
// use file sizes.
func (s *Statter) DiskUsageSimple(p Prefix, path string) (*Result, error) {
	return s.DiskUsage(p, path)
}
