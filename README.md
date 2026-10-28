# clistat

A Go library for measuring and reporting system resource usage.

## Overview

`clistat` provides utilities to query resource usage metrics on a system including:

- Host resources:
  - CPU usage
  - Memory usage
  - Disk usage

- Container resources (cgroup v1 and v2 support):
  - CPU usage
  - Memory usage

## Features

- Human-readable output formatting with appropriate unit prefixes
- Support for both host and containerized environments
- Automatic container detection
- Support for both cgroup v1 and v2
- Configurable with options pattern

## Usage

The below sample code can be run as follows:

```shell
CGO_ENABLED=0 go run github.com/coder/clistat/cmd/clistat@main
```

```go
package main

import (
	"fmt"
	"log"

	"github.com/coder/clistat"
)

func main() {
	// Create a new Statter
	s, err := clistat.New()
	if err != nil {
		log.Fatal(err)
	}

	// Query host CPU usage
	cpu, err := s.HostCPU()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("CPU: %s\n", cpu)

	// Query host memory usage
	mem, err := s.HostMemory(clistat.PrefixGibi)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Memory: %s\n", mem)

	// Query disk usage
	disk, err := s.Disk(clistat.PrefixGibi, "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Disk: %s\n", disk)

	// Check if running in a container
	isContainer, err := s.IsContainerized()
	if err != nil {
		log.Fatal(err)
	}

	if isContainer {
		// Query container CPU usage
		containerCPU, err := s.ContainerCPU()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Container CPU: %s\n", containerCPU)

		// Query container memory usage
		containerMem, err := s.ContainerMemory(clistat.PrefixMebi)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Container Memory: %s\n", containerMem)
	}
}
```

## Output Examples

The library formats output in a human-friendly way:

```
CPU: 2.5/8 cores (31%)
Memory: 3.2/16 GiB (20%)
Disk: 50/200 GiB (25%)
Container CPU: 0.5/2.5 cores (20%)
Container Memory: 256/1024 MiB (25%)
```

## How container limits are reported

On cgroup v2, a container's cgroup can have child cgroups. For example, in a
docker-in-docker workspace the process runs in `/init.scope` and nested
containers run in `/docker/<id>`. Limits can be set at any level, and the
kernel enforces every limit from the process's cgroup up to the container root.

`ContainerCPU` and `ContainerMemory` handle this as follows:

1. Check each cgroup from the process's cgroup up to the container root for a
   limit (`cpu.max` for CPU, `memory.max` for memory).
1. For each limit found, read usage from the same cgroup that sets it. This
   usage includes every cgroup below it, so processes that share the limit
   are counted.
   - CPU: the change in `usage_usec` from `cpu.stat` over the sample interval.
   - Memory: `memory.current` minus `inactive_file` from `memory.stat`.
1. Report the limit with the least room left (limit minus usage), since that
   is the one that will be hit first. If two limits have the same room left,
   the one nearest to the process is reported.
1. If no cgroup sets a limit, report usage of the process's own cgroup with
   no total.

If a cgroup sets a limit but does not have the usage files, usage of the
process's own cgroup is used for that limit.

Only one limit is reported. The result answers "how close is this process to
being throttled or killed," not "what are all the limits."

### Limits outside the container

clistat can only see cgroups inside the container's cgroup namespace. Limits
set above the container root are not visible. For example, a Kubernetes
pod-level limit is not visible from inside a container that has its own
cgroup namespace. If such a limit is tighter than the container's own limit,
clistat reports the container's limit.

On cgroup v1, clistat reads the limit and usage from the container's own
cgroup only.

## Testing

The library includes a comprehensive test suite that can be run with:

```
go test ./...
```

Some tests are designed to detect containerization status and will adjust based on the environment. You can control test behavior with these environment variables:

- `CLISTAT_IS_CONTAINERIZED=yes` - Indicate tests are running in a container
- `CLISTAT_HAS_MEMORY_LIMIT=yes` - Indicate container has memory limits
- `CLISTAT_HAS_CPU_LIMIT=yes` - Indicate container has CPU limits
