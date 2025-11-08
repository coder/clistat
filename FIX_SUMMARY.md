# CgroupV2 Memory Parent Lookup Fix

## Problem Identified

The `memory()` function in `cgroupv2.go` had a critical flaw where it did not implement parent lookup logic for `memory.current` and `memory.stat` files when they were missing in the current cgroup.

### Affected Files

- **`memory.current`** - Current memory usage in bytes
- **`memory.stat`** - Memory statistics including `inactive_file`

### Inconsistent Behavior

While `memory.max` correctly implemented parent lookup via the `memoryMaxBytes()` helper method, the other two memory files directly called low-level read functions and immediately returned errors if files didn't exist.

## Root Cause

In the original `memory()` function (lines 153-180):

```go
// ❌ Direct read - no parent fallback
currUsageBytes, err := readInt64(s.fs, memoryUsagePath)
if err != nil {
    return nil, xerrors.Errorf("read memory usage: %w", err)
}

// ❌ Direct read - no parent fallback
inactiveFileBytes, err := readInt64Prefix(s.fs, memoryStatPath, "inactive_file")
if err != nil {
    return nil, xerrors.Errorf("read memory stats: %w", err)
}
```

This violated the user's requirement that "logic should lookup the parent if not found for the current group."

## Solution Implemented

Added two new helper methods following the same pattern as `memoryMaxBytes()`:

### 1. `memoryCurrentBytes()` (lines 153-177)

```go
func (s cgroupV2Statter) memoryCurrentBytes() (int64, error) {
    memoryUsagePath := filepath.Join(s.path, cgroupV2MemoryUsageBytes)

    currUsageBytes, err := readInt64(s.fs, memoryUsagePath)
    if err != nil {
        if !errors.Is(err, fs.ErrNotExist) {
            return 0, xerrors.Errorf("read memory current: %w", err)
        }

        // Parent fallback logic
        if s.parent != nil {
            result, err := s.parent.memoryCurrentBytes()
            if err != nil {
                return 0, xerrors.Errorf("read parent memory current: %w", err)
            }
            return result, nil
        }

        return 0, xerrors.Errorf("read memory current: %w", err)
    }

    return currUsageBytes, nil
}
```

### 2. `memoryInactiveFileBytes()` (lines 179-203)

```go
func (s cgroupV2Statter) memoryInactiveFileBytes() (int64, error) {
    memoryStatPath := filepath.Join(s.path, cgroupV2MemoryStat)

    inactiveFileBytes, err := readInt64Prefix(s.fs, memoryStatPath, "inactive_file")
    if err != nil {
        if !errors.Is(err, fs.ErrNotExist) {
            return 0, xerrors.Errorf("read memory stat inactive_file: %w", err)
        }

        // Parent fallback logic
        if s.parent != nil {
            result, err := s.parent.memoryInactiveFileBytes()
            if err != nil {
                return 0, xerrors.Errorf("read parent memory stat inactive_file: %w", err)
            }
            return result, nil
        }

        return 0, xerrors.Errorf("read memory stat inactive_file: %w", err)
    }

    return inactiveFileBytes, nil
}
```

### 3. Updated `memory()` function (lines 205-229)

```go
func (s cgroupV2Statter) memory(p Prefix) (*Result, error) {
    // ... setup code ...

    // ✅ Now uses helper with parent fallback
    currUsageBytes, err := s.memoryCurrentBytes()
    if err != nil {
        return nil, xerrors.Errorf("read memory usage: %w", err)
    }

    // ✅ Now uses helper with parent fallback
    inactiveFileBytes, err := s.memoryInactiveFileBytes()
    if err != nil {
        return nil, xerrors.Errorf("read memory stats: %w", err)
    }

    r.Used = float64(currUsageBytes - inactiveFileBytes)
    return r, nil
}
```

## Key Design Decisions

1. **Only fallback on `fs.ErrNotExist`**: Other errors (parsing, permissions) are immediately returned
2. **Recursive parent lookup**: Follows the same pattern as `memoryMaxBytes()`
3. **Consistent error messages**: Uses similar error wrapping throughout
4. **Protected by depth limit**: Existing `maxSupportedCgroupDepth` prevents infinite recursion

## Test Coverage

### Existing Tests (All Pass)
- ✅ `TestRecursiveCreation` - Validates parent chain creation
- ✅ `TestStatter/CgroupV2/ContainerMemory/Limit` - Standard memory with limits
- ✅ `TestStatter/CgroupV2/ContainerMemory/NoLimit` - Memory with no limits
- ✅ `TestStatter/CgroupV2/Kubernetes/Memory/LimitInParent` - Parent limit lookup
- ✅ `TestStatter/CgroupV2/Kubernetes/Memory/NoLimit` - No limits at any level

### New Tests Added
- ✅ `TestMemoryCurrentParentLookup` - Verifies `memory.current` falls back to parent
- ✅ `TestMemoryStatParentLookup` - Verifies `memory.stat` falls back to parent

## Impact

### Before Fix
❌ Child cgroups without `memory.current` or `memory.stat` would fail with errors  
❌ Inconsistent with cgroupv2 semantics where values can be inherited  
❌ Different behavior from `memory.max` which worked correctly  

### After Fix
✅ Child cgroups correctly inherit values from parents when files are missing  
✅ Consistent behavior across all memory-related file reads  
✅ Matches user requirement for parent lookup  
✅ All tests pass including new parent lookup tests  

## Files Modified

- `cgroupv2.go` - Added helper methods and updated `memory()` function
- `cgroupv2_memory_parent_test.go` - New test file demonstrating the fix

## Verification

```bash
# All cgroupv2 tests pass
go test -v -run "TestStatter/CgroupV2"
# PASS: All 12 cgroupv2 test cases

# New parent lookup tests pass
go test -v -run TestMemoryCurrentParentLookup
go test -v -run TestMemoryStatParentLookup
# PASS: Both tests

# Build succeeds
go build ./...
# Success
```

## Conclusion

The fix ensures that `memory.current` and `memory.stat` now have the same robust parent lookup behavior as `memory.max`, resolving the inconsistency and meeting the user's requirement that "logic should lookup the parent if not found for the current group."
