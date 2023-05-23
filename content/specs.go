package content

import (
	"fmt"
	"strconv"

	system "github.com/mr-cheeezz/GoPC/core"
)

type SpecGetterFunc func() string

type Monitor struct {
	ID          string
	Resolution  string
	RefreshRate string
}

var Specs = map[string]SpecGetterFunc{
	"CPU":         system.Cpu,
	"GPU":         system.Gpu,
	"RAM":         system.Memory,
	"Motherboard": system.MotherBoard,
}

var Physical = map[string][]SpecGetterFunc{
	"Monitor": MonitorInfo(),
}

var SoftW = map[string]SpecGetterFunc{
	"OS":     system.OSName,
	"Kernel": system.KernelVer,
}

func Monitors() string {
	numMonitors, err := strconv.Atoi(system.MonitorNum())
	if err != nil {
		fmt.Println("Error converting number of monitors to integer:", err)
		return ""
	}

	if numMonitors > 1 {
		return "Monitors"
	} else {
		return "Monitor"
	}
}

func MonitorInfo() []SpecGetterFunc {
	numMonitors, err := strconv.Atoi(system.MonitorNum())
	if err != nil {
		fmt.Println("Error converting number of monitors to integer:", err)
		return nil
	}

	monitorFuncs := make([]SpecGetterFunc, numMonitors)
	for i := 0; i < numMonitors; i++ {
		monitorFuncs[i] = func() string {
			return fmt.Sprintf("%s, %sHz", system.GetMonitorResolution(i), system.GetMonitorRefreshRate(i))
		}
	}

	return monitorFuncs
}
