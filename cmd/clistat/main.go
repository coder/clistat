package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/coder/clistat"
)

func main() {
	// Create a new Statter
	s, err := clistat.New()
	if err != nil {
		log.Fatal(err)
	}

	// Query host CPU usage
	if runtime.GOOS == "darwin" {
		fmt.Printf("CPU: unsupported on macOS\n")
	} else {
		cpu, err := s.HostCPU()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("CPU: %s\n", cpu)
	}

	// Query host memory usage
	mem, err := s.HostMemory(clistat.PrefixGibi)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Memory: %s\n", mem)

	// Query disk usage
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	disk, err := s.Disk(clistat.PrefixGibi, wd)
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
