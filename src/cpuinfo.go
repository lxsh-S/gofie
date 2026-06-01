package src

import (
	"github.com/shirou/gopsutil/v4/cpu"
)

func GetCPU() string {
	// First lets fetch

	cpuinfo, err := cpu.Info()
	if err != nil {
		return "Failed to fetch cpu Info..."
	}
	CpuName := cpuinfo[0].ModelName
	return CpuName
}
