package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	fmt.Println("Operating System:", runtime.GOOS)
	fmt.Println("Architecture:", runtime.GOARCH)

	hostname, err := os.Hostname()
	if err == nil {
	}

	username := os.Getenv("USERNAME")
	if username == "" {
		username = os.Getenv("USER")
	}
	fmt.Println(fmt.Sprintf("%s@%s", hostname, username))

	cmd := exec.Command("wmic", "cpu", "get", "Name")
	output, err := cmd.Output()
	if err == nil {
		lines := strings.Split(string(output), "\n")
		if len(lines) >= 2 {
			cpuInfo := strings.TrimSpace(lines[1])
			fmt.Println("CPU:", cpuInfo)
		}
	}

	cmd = exec.Command("wmic", "memorychip", "get", "Capacity")
	output, err = cmd.Output()
	if err == nil {
		lines := strings.Split(string(output), "\n")
		totalMemory := int64(0)
		for _, line := range lines[1:] {
			if strings.TrimSpace(line) != "" {
				var capacity int64
				_, err := fmt.Sscan(strings.TrimSpace(line), &capacity)
				if err == nil {
					totalMemory += capacity
				}
			}
		}
		fmt.Println("Memory:", totalMemory/1024/1024, "MB")
	}
}
