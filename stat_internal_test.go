package clistat

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"tailscale.com/types/ptr"
)

func TestResultString(t *testing.T) {
	t.Parallel()
	for _, tt := range []struct {
		Expected string
		Result   Result
	}{
		{
			Expected: "1.23/5.68 quatloos (22%)",
			Result:   Result{Used: 1.234, Total: ptr.To(5.678), Unit: "quatloos"},
		},
		{
			Expected: "0/0 HP",
			Result:   Result{Used: 0.0, Total: ptr.To(0.0), Unit: "HP"},
		},
		{
			Expected: "123 seconds",
			Result:   Result{Used: 123.01, Total: nil, Unit: "seconds"},
		},
		{
			Expected: "12.3",
			Result:   Result{Used: 12.34, Total: nil, Unit: ""},
		},
		{
			Expected: "1.5 KiB",
			Result:   Result{Used: 1536, Total: nil, Unit: "B", Prefix: PrefixKibi},
		},
		{
			Expected: "1.23 things",
			Result:   Result{Used: 1.234, Total: nil, Unit: "things"},
		},
		{
			Expected: "0/100 TiB (0%)",
			Result:   Result{Used: 1, Total: ptr.To(100.0 * float64(PrefixTebi)), Unit: "B", Prefix: PrefixTebi},
		},
		{
			Expected: "0.5/8 cores (6%)",
			Result:   Result{Used: 0.5, Total: ptr.To(8.0), Unit: "cores"},
		},
	} {
		assert.Equal(t, tt.Expected, tt.Result.String())
	}
}

func TestStatter(t *testing.T) {
	t.Parallel()

	// We cannot make many assertions about the data we get back
	// for host-specific measurements because these tests could
	// and should run successfully on any OS.
	// The best we can do is assert that it is non-zero.
	t.Run("HostOnly", func(t *testing.T) {
		t.Parallel()
		fs := initFS(t, fsHostOnly)
		s, err := New(WithFS(fs))
		require.NoError(t, err)
		t.Run("HostCPU", func(t *testing.T) {
			t.Parallel()
			cpu, err := s.HostCPU()
			require.NoError(t, err)
			// assert.NotZero(t, cpu.Used) // HostCPU can sometimes be zero.
			assert.NotZero(t, cpu.Total)
			assert.Equal(t, "cores", cpu.Unit)
		})

		t.Run("HostMemory", func(t *testing.T) {
			t.Parallel()
			mem, err := s.HostMemory(PrefixDefault)
			require.NoError(t, err)
			assert.NotZero(t, mem.Used)
			assert.NotZero(t, mem.Total)
			assert.Equal(t, "B", mem.Unit)
		})

		t.Run("HostDisk", func(t *testing.T) {
			t.Parallel()
			disk, err := s.Disk(PrefixDefault, "") // default to home dir
			require.NoError(t, err)
			assert.NotZero(t, disk.Used)
			assert.NotZero(t, disk.Total)
			assert.Equal(t, "B", disk.Unit)
		})
	})

	t.Run("RealMemory", func(t *testing.T) {
		t.Parallel()

		fs := afero.NewOsFs()
		s, err := New(WithFS(fs))
		require.NoError(t, err)

		mem, err := s.HostMemory(PrefixDefault)
		require.NoError(t, err)
		assert.NotNil(t, mem)
		assert.NotZero(t, mem.Used, "Memory usage should be non-zero")
		assert.NotZero(t, mem.Total, "Total memory should be non-zero")
		assert.True(t, *mem.Total > mem.Used, "Total memory should be greater than used memory")
	})

	// Sometimes we do need to "fake" some stuff
	// that happens while we wait.
	withWait := func(waitF func(time.Duration)) Option {
		return func(s *Statter) {
			s.wait = waitF
		}
	}

	// Other times we just want things to run fast.
	withNoWait := func(s *Statter) {
		s.wait = func(time.Duration) {}
	}

	// We don't want to use the actual host CPU here.
	withNproc := func(n int) Option {
		return func(s *Statter) {
			s.nproc = n
		}
	}

	withIsCGroupV2 := func(state bool) Option {
		return func(s *Statter) {
			s.cgroupV2Detector = func(_ afero.Fs) bool {
				return state
			}
		}
	}

	// For container-specific measurements, everything we need
	// can be read from the filesystem. We control the FS, so
	// we control the data.
	t.Run("CGroupV1", func(t *testing.T) {
		t.Parallel()

		t.Run("ContainerCPU/Limit", func(t *testing.T) {
			t.Parallel()

			fs := initFS(t, fsContainerCgroupV1)
			fakeWait := func(time.Duration) {
				// Fake 1 second in ns of usage
				mungeFS(t, fs, cgroupV1CPUAcctUsage, "100000000")
			}
			s, err := New(WithFS(fs), withWait(fakeWait), withIsCGroupV2(false))
			require.NoError(t, err)

			cpu, err := s.ContainerCPU()
			require.NoError(t, err)

			require.NotNil(t, cpu)
			assert.Equal(t, 1.0, cpu.Used)
			require.NotNil(t, cpu.Total)
			assert.Equal(t, 2.5, *cpu.Total)
			assert.Equal(t, "cores", cpu.Unit)
		})

		t.Run("ContainerCPU/NoLimit", func(t *testing.T) {
			t.Parallel()

			fs := initFS(t, fsContainerCgroupV1NoLimit)
			fakeWait := func(time.Duration) {
				// Fake 1 second in ns of usage
				mungeFS(t, fs, cgroupV1CPUAcctUsage, "100000000")
			}
			s, err := New(WithFS(fs), withNproc(2), withWait(fakeWait), withIsCGroupV2(false))
			require.NoError(t, err)

			cpu, err := s.ContainerCPU()
			require.NoError(t, err)

			require.NotNil(t, cpu)
			assert.Equal(t, 1.0, cpu.Used)
			require.Nil(t, cpu.Total)
			assert.Equal(t, "cores", cpu.Unit)
		})

		t.Run("ContainerCPU/AltPath", func(t *testing.T) {
			t.Parallel()

			fs := initFS(t, fsContainerCgroupV1AltPath)
			fakeWait := func(time.Duration) {
				// Fake 1 second in ns of usage
				mungeFS(t, fs, "/sys/fs/cgroup/cpuacct/cpuacct.usage", "100000000")
			}
			s, err := New(WithFS(fs), withNproc(2), withWait(fakeWait), withIsCGroupV2(false))
			require.NoError(t, err)

			cpu, err := s.ContainerCPU()
			require.NoError(t, err)

			require.NotNil(t, cpu)
			assert.Equal(t, 1.0, cpu.Used)
			require.NotNil(t, cpu.Total)
			assert.Equal(t, 2.5, *cpu.Total)
			assert.Equal(t, "cores", cpu.Unit)
		})

		t.Run("ContainerMemory", func(t *testing.T) {
			t.Parallel()

			fs := initFS(t, fsContainerCgroupV1)
			s, err := New(WithFS(fs), withNoWait, withIsCGroupV2(false))
			require.NoError(t, err)

			mem, err := s.ContainerMemory(PrefixDefault)
			require.NoError(t, err)

			require.NotNil(t, mem)
			assert.Equal(t, 268435456.0, mem.Used)
			assert.NotNil(t, mem.Total)
			assert.Equal(t, 1073741824.0, *mem.Total)
			assert.Equal(t, "B", mem.Unit)
		})

		t.Run("ContainerMemory/NoLimit", func(t *testing.T) {
			t.Parallel()

			fs := initFS(t, fsContainerCgroupV1NoLimit)
			s, err := New(WithFS(fs), withNoWait, withIsCGroupV2(false))
			require.NoError(t, err)

			mem, err := s.ContainerMemory(PrefixDefault)
			require.NoError(t, err)

			require.NotNil(t, mem)
			assert.Equal(t, 268435456.0, mem.Used)
			assert.Nil(t, mem.Total)
			assert.Equal(t, "B", mem.Unit)
		})

		t.Run("ContainerMemory/NoLimit", func(t *testing.T) {
			t.Parallel()

			fs := initFS(t, fsContainerCgroupV1DockerNoMemoryLimit)
			s, err := New(WithFS(fs), withNoWait, withIsCGroupV2(false))
			require.NoError(t, err)

			mem, err := s.ContainerMemory(PrefixDefault)
			require.NoError(t, err)

			require.NotNil(t, mem)
			assert.Equal(t, 268435456.0, mem.Used)
			assert.Nil(t, mem.Total)
			assert.Equal(t, "B", mem.Unit)
		})
	})

	t.Run("CGroupV2", func(t *testing.T) {
		t.Parallel()

		t.Run("ContainerCPU/Limit", func(t *testing.T) {
			t.Parallel()

			fs := initFS(t, fsContainerCgroupV2)
			fakeWait := func(time.Duration) {
				mungeFS(t, fs, cgroupV2CPUStat, "usage_usec 100000")
			}
			s, err := New(WithFS(fs), withWait(fakeWait), withIsCGroupV2(true))

			require.NoError(t, err)
			cpu, err := s.ContainerCPU()
			require.NoError(t, err)

			require.NotNil(t, cpu)
			assert.Equal(t, 1.0, cpu.Used)
			require.NotNil(t, cpu.Total)
			assert.Equal(t, 2.5, *cpu.Total)
			assert.Equal(t, "cores", cpu.Unit)
		})

		t.Run("ContainerCPU/NoLimit", func(t *testing.T) {
			t.Parallel()

			fs := initFS(t, fsContainerCgroupV2NoLimit)
			fakeWait := func(time.Duration) {
				mungeFS(t, fs, cgroupV2CPUStat, "usage_usec 100000")
			}
			s, err := New(WithFS(fs), withNproc(2), withWait(fakeWait), withIsCGroupV2(true))
			require.NoError(t, err)

			cpu, err := s.ContainerCPU()
			require.NoError(t, err)

			require.NotNil(t, cpu)
			assert.Equal(t, 1.0, cpu.Used)
			require.Nil(t, cpu.Total)
			assert.Equal(t, "cores", cpu.Unit)
		})

		t.Run("ContainerCPU/MissingCPUMax", func(t *testing.T) {
			t.Parallel()

			fs := initFS(t, fsContainerCgroupV2NoCPUMax)
			fakeWait := func(time.Duration) {
				mungeFS(t, fs, cgroupV2CPUStat, "usage_usec 100000")
			}
			s, err := New(WithFS(fs), withNproc(2), withWait(fakeWait), withIsCGroupV2(true))
			require.NoError(t, err)

			cpu, err := s.ContainerCPU()
			require.NoError(t, err)

			require.NotNil(t, cpu)
			assert.Equal(t, 1.0, cpu.Used)
			require.Nil(t, cpu.Total)
			assert.Equal(t, "cores", cpu.Unit)
		})

		t.Run("ContainerMemory/Limit", func(t *testing.T) {
			t.Parallel()

			fs := initFS(t, fsContainerCgroupV2)
			s, err := New(WithFS(fs), withNoWait, withIsCGroupV2(true))
			require.NoError(t, err)

			mem, err := s.ContainerMemory(PrefixDefault)
			require.NoError(t, err)

			require.NotNil(t, mem)
			assert.Equal(t, 268435456.0, mem.Used)
			assert.NotNil(t, mem.Total)
			assert.Equal(t, 1073741824.0, *mem.Total)
			assert.Equal(t, "B", mem.Unit)
		})

		t.Run("ContainerMemory/NoLimit", func(t *testing.T) {
			t.Parallel()

			fs := initFS(t, fsContainerCgroupV2NoLimit)
			s, err := New(WithFS(fs), withNoWait, withIsCGroupV2(true))
			require.NoError(t, err)

			mem, err := s.ContainerMemory(PrefixDefault)
			require.NoError(t, err)

			require.NotNil(t, mem)
			assert.Equal(t, 268435456.0, mem.Used)
			assert.Nil(t, mem.Total)
			assert.Equal(t, "B", mem.Unit)
		})

		t.Run("ContainerMemory/MissingLimit", func(t *testing.T) {
			t.Parallel()

			fs := initFS(t, fsContainerCgroupV2NoMemoryMax)
			s, err := New(WithFS(fs), withNoWait, withIsCGroupV2(true))
			require.NoError(t, err)

			mem, err := s.ContainerMemory(PrefixDefault)
			require.NoError(t, err)

			require.NotNil(t, mem)
			assert.Equal(t, 268435456.0, mem.Used)
			assert.Nil(t, mem.Total)
			assert.Equal(t, "B", mem.Unit)
		})
	})
}

func TestCGroupV2Detection(t *testing.T) {
	t.Parallel()

	hostISCGroupV2 := os.Getenv("CLISTAT_IS_CGROUPV2") == "yes"

	tests := []struct {
		name string
		fs   afero.Fs
	}{
		{
			name: "OsFs",
			fs:   afero.NewOsFs(),
		},
		{
			name: "OsFs/ReadOnly",
			fs:   afero.NewReadOnlyFs(afero.NewOsFs()),
		},
		{
			name: "InMemoryFs/MissingCPUMax",
			fs:   initFS(t, fsContainerCgroupV2NoCPUMax),
		},
		{
			name: "InMemoryFs/MissingMemoryMax",
			fs:   initFS(t, fsContainerCgroupV2NoMemoryMax),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s, err := New(WithFS(tt.fs))
			require.NoError(t, err)

			isCGroupV2 := s.cgroupV2Detector(s.fs)
			assert.Equal(t, hostISCGroupV2, isCGroupV2)
		})
	}
}

func TestIsContainerized(t *testing.T) {
	t.Parallel()

	t.Run("IsCorrectlyDetected", func(t *testing.T) {
		t.Parallel()

		s, err := New(WithFS(afero.NewOsFs()))
		require.NoError(t, err)
		isContainer, err := s.IsContainerized()
		require.NoError(t, err)

		if os.Getenv("CLISTAT_IS_CONTAINERIZED") == "yes" {
			assert.True(t, isContainer, "Expected to be detected as running in a container")
		} else {
			assert.False(t, isContainer, "Expected to be detected as not running in a container")
		}
	})

	t.Run("InContainer", func(t *testing.T) {
		t.Parallel()

		if os.Getenv("CLISTAT_IS_CONTAINERIZED") != "yes" {
			t.Skip("Skipping test - CLISTAT_IS_CONTAINERIZED is not set to 'yes'")
		}

		t.Run("ContainerMemory", func(t *testing.T) {
			t.Parallel()

			fs := afero.NewOsFs()
			s, err := New(WithFS(fs))
			require.NoError(t, err)

			mem, err := s.ContainerMemory(PrefixDefault)
			require.NoError(t, err)
			require.NotNil(t, mem)
			assert.NotZero(t, mem.Used, "Container memory usage should be non-zero")

			hasMemoryLimit := os.Getenv("CLISTAT_HAS_MEMORY_LIMIT") == "yes"
			if hasMemoryLimit {
				require.NotNil(t, mem.Total, "Container should have memory limit")
				assert.NotZero(t, *mem.Total, "Container total memory should be non-zero")
				assert.True(t, *mem.Total > mem.Used, "Container total memory should be greater than used memory")
			} else {
				assert.Nil(t, mem.Total, "Container should not have memory limit")
			}
		})

		t.Run("ContainerCPU", func(t *testing.T) {
			t.Parallel()

			fs := afero.NewOsFs()
			s, err := New(WithFS(fs))
			require.NoError(t, err)

			cpu, err := s.ContainerCPU()
			require.NoError(t, err)
			require.NotNil(t, cpu)
			assert.NotZero(t, cpu.Used, "Container CPU usage should be non-zero")

			hasCPULimit := os.Getenv("CLISTAT_HAS_CPU_LIMIT") == "yes"
			if hasCPULimit {
				require.NotNil(t, cpu.Total, "Container should have CPU limit")
				assert.NotZero(t, *cpu.Total, "Container total CPU should be non-zero")
				assert.True(t, *cpu.Total > cpu.Used, "Container total CPU should be greater than used CPU")
			} else {
				assert.Nil(t, cpu.Total, "Container should not have CPU limit")
			}
		})
	})

	for _, tt := range []struct {
		Name     string
		FS       map[string]string
		Expected bool
		Error    string
	}{
		{
			Name:     "Empty",
			FS:       map[string]string{},
			Expected: false,
			Error:    "",
		},
		{
			Name:     "BareMetal",
			FS:       fsHostOnly,
			Expected: false,
			Error:    "",
		},
		{
			Name:     "Docker",
			FS:       fsContainerCgroupV1,
			Expected: true,
			Error:    "",
		},
		{
			Name:     "Sysbox",
			FS:       fsContainerSysbox,
			Expected: true,
			Error:    "",
		},
		{
			Name:     "Docker (Cgroupns=private)",
			FS:       fsContainerCgroupV2PrivateCgroupns,
			Expected: true,
			Error:    "",
		},
	} {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			fs := initFS(t, tt.FS)
			s, err := New(WithFS(fs))
			require.NoError(t, err)

			actual, err := s.IsContainerized()
			if tt.Error == "" {
				assert.NoError(t, err)
				assert.Equal(t, tt.Expected, actual)
			} else {
				assert.ErrorContains(t, err, tt.Error)
				assert.False(t, actual)
			}
		})
	}
}

func TestCPUTotal(t *testing.T) {
	t.Parallel()

	expectedCPUCountString := os.Getenv("CLISTAT_CPU_COUNT")
	if expectedCPUCountString == "" {
		t.Skip("Skipping test - CLISTAT_CPU_COUNT not set")
	}

	expectedCPUCount, err := strconv.Atoi(expectedCPUCountString)
	require.NoError(t, err)

	s, err := New()
	require.NoError(t, err)

	cpu, err := s.HostCPU()
	require.NoError(t, err)
	require.NotNil(t, cpu.Total)
	require.Equal(t, expectedCPUCount, int(*cpu.Total))
}

// helper function for initializing a fs
func initFS(t testing.TB, m map[string]string) afero.Fs {
	t.Helper()
	fs := afero.NewMemMapFs()
	for k, v := range m {
		mungeFS(t, fs, k, v)
	}
	return fs
}

// helper function for writing v to fs under path k
func mungeFS(t testing.TB, fs afero.Fs, k, v string) {
	t.Helper()
	require.NoError(t, afero.WriteFile(fs, k, []byte(v+"\n"), 0o600))
}

var (
	fsHostOnly = map[string]string{
		procOneCgroup: "0::/",
		procMounts:    "/dev/sda1 / ext4 rw,relatime 0 0",
	}
	fsContainerSysbox = map[string]string{
		procOneCgroup: "0::/docker/aa86ac98959eeedeae0ecb6e0c9ddd8ae8b97a9d0fdccccf7ea7a474f4e0bb1f",
		procMounts: `overlay / overlay rw,relatime,lowerdir=/some/path:/some/path,upperdir=/some/path:/some/path,workdir=/some/path:/some/path 0 0
sysboxfs /proc/sys proc ro,nosuid,nodev,noexec,relatime 0 0`,
		cgroupV2CPUMax:  "250000 100000",
		cgroupV2CPUStat: "usage_usec 0",
	}
	fsContainerCgroupV2 = map[string]string{
		procOneCgroup: "0::/docker/aa86ac98959eeedeae0ecb6e0c9ddd8ae8b97a9d0fdccccf7ea7a474f4e0bb1f",
		procMounts: `overlay / overlay rw,relatime,lowerdir=/some/path:/some/path,upperdir=/some/path:/some/path,workdir=/some/path:/some/path 0 0
proc /proc/sys proc ro,nosuid,nodev,noexec,relatime 0 0`,
		cgroupV2CPUMax:           "250000 100000",
		cgroupV2CPUStat:          "usage_usec 0",
		cgroupV2MemoryMaxBytes:   "1073741824",
		cgroupV2MemoryUsageBytes: "536870912",
		cgroupV2MemoryStat:       "inactive_file 268435456",
	}
	fsContainerCgroupV2NoLimit = map[string]string{
		procOneCgroup: "0::/docker/aa86ac98959eeedeae0ecb6e0c9ddd8ae8b97a9d0fdccccf7ea7a474f4e0bb1f",
		procMounts: `overlay / overlay rw,relatime,lowerdir=/some/path:/some/path,upperdir=/some/path:/some/path,workdir=/some/path:/some/path 0 0
proc /proc/sys proc ro,nosuid,nodev,noexec,relatime 0 0`,
		cgroupV2CPUMax:           "max 100000",
		cgroupV2CPUStat:          "usage_usec 0",
		cgroupV2MemoryMaxBytes:   "max",
		cgroupV2MemoryUsageBytes: "536870912",
		cgroupV2MemoryStat:       "inactive_file 268435456",
	}
	fsContainerCgroupV2PrivateCgroupns = map[string]string{
		procOneCgroup: "0::/",
		procMounts: `overlay / overlay rw,relatime,lowerdir=/some/path:/some/path,upperdir=/some/path:/some/path,workdir=/some/path:/some/path 0 0
proc /proc/sys proc ro,nosuid,nodev,noexec,relatime 0 0`,
		sysCgroupType: "domain",
	}
	fsContainerCgroupV2NoCPUMax = map[string]string{
		procOneCgroup: "0::/docker/aa86ac98959eeedeae0ecb6e0c9ddd8ae8b97a9d0fdccccf7ea7a474f4e0bb1f",
		procMounts: `overlay / overlay rw,relatime,lowerdir=/some/path:/some/path,upperdir=/some/path:/some/path,workdir=/some/path:/some/path 0 0
proc /proc/sys proc ro,nosuid,nodev,noexec,relatime 0 0`,
		// Note: cpu.max is intentionally missing
		cgroupV2CPUStat:          "usage_usec 0",
		cgroupV2MemoryMaxBytes:   "1073741824",
		cgroupV2MemoryUsageBytes: "536870912",
		cgroupV2MemoryStat:       "inactive_file 268435456",
	}
	fsContainerCgroupV2NoMemoryMax = map[string]string{
		procOneCgroup: "0::/docker/aa86ac98959eeedeae0ecb6e0c9ddd8ae8b97a9d0fdccccf7ea7a474f4e0bb1f",
		procMounts: `overlay / overlay rw,relatime,lowerdir=/some/path:/some/path,upperdir=/some/path:/some/path,workdir=/some/path:/some/path 0 0
proc /proc/sys proc ro,nosuid,nodev,noexec,relatime 0 0`,
		cgroupV2CPUMax:           "max 100000",
		cgroupV2CPUStat:          "usage_usec 0",
		cgroupV2MemoryUsageBytes: "536870912",
		cgroupV2MemoryStat:       "inactive_file 268435456",
	}
	fsContainerCgroupV1 = map[string]string{
		procOneCgroup: "0::/docker/aa86ac98959eeedeae0ecb6e0c9ddd8ae8b97a9d0fdccccf7ea7a474f4e0bb1f",
		procMounts: `overlay / overlay rw,relatime,lowerdir=/some/path:/some/path,upperdir=/some/path:/some/path,workdir=/some/path:/some/path 0 0
proc /proc/sys proc ro,nosuid,nodev,noexec,relatime 0 0`,
		cgroupV1CPUAcctUsage:        "0",
		cgroupV1CFSQuotaUs:          "250000",
		cgroupV1CFSPeriodUs:         "100000",
		cgroupV1MemoryMaxUsageBytes: "1073741824",
		cgroupV1MemoryUsageBytes:    "536870912",
		cgroupV1MemoryStat:          "total_inactive_file 268435456",
	}
	fsContainerCgroupV1NoLimit = map[string]string{
		procOneCgroup: "0::/docker/aa86ac98959eeedeae0ecb6e0c9ddd8ae8b97a9d0fdccccf7ea7a474f4e0bb1f",
		procMounts: `overlay / overlay rw,relatime,lowerdir=/some/path:/some/path,upperdir=/some/path:/some/path,workdir=/some/path:/some/path 0 0
proc /proc/sys proc ro,nosuid,nodev,noexec,relatime 0 0`,
		cgroupV1CPUAcctUsage:        "0",
		cgroupV1CFSQuotaUs:          "-1",
		cgroupV1CFSPeriodUs:         "100000",
		cgroupV1MemoryMaxUsageBytes: "max", // I have never seen this in the wild
		cgroupV1MemoryUsageBytes:    "536870912",
		cgroupV1MemoryStat:          "total_inactive_file 268435456",
	}
	fsContainerCgroupV1DockerNoMemoryLimit = map[string]string{
		procOneCgroup: "0::/docker/aa86ac98959eeedeae0ecb6e0c9ddd8ae8b97a9d0fdccccf7ea7a474f4e0bb1f",
		procMounts: `overlay / overlay rw,relatime,lowerdir=/some/path:/some/path,upperdir=/some/path:/some/path,workdir=/some/path:/some/path 0 0
proc /proc/sys proc ro,nosuid,nodev,noexec,relatime 0 0`,
		cgroupV1CPUAcctUsage:        "0",
		cgroupV1CFSQuotaUs:          "-1",
		cgroupV1CFSPeriodUs:         "100000",
		cgroupV1MemoryMaxUsageBytes: "9223372036854771712",
		cgroupV1MemoryUsageBytes:    "536870912",
		cgroupV1MemoryStat:          "total_inactive_file 268435456",
	}
	fsContainerCgroupV1AltPath = map[string]string{
		procOneCgroup: "0::/docker/aa86ac98959eeedeae0ecb6e0c9ddd8ae8b97a9d0fdccccf7ea7a474f4e0bb1f",
		procMounts: `overlay / overlay rw,relatime,lowerdir=/some/path:/some/path,upperdir=/some/path:/some/path,workdir=/some/path:/some/path 0 0
proc /proc/sys proc ro,nosuid,nodev,noexec,relatime 0 0`,
		"/sys/fs/cgroup/cpuacct/cpuacct.usage": "0",
		"/sys/fs/cgroup/cpu/cpu.cfs_quota_us":  "250000",
		"/sys/fs/cgroup/cpu/cpu.cfs_period_us": "100000",
		cgroupV1MemoryMaxUsageBytes:            "1073741824",
		cgroupV1MemoryUsageBytes:               "536870912",
		cgroupV1MemoryStat:                     "total_inactive_file 268435456",
	}
)
