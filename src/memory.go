package src

import (
	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/v4/mem"
)

func GETMemoryUsed() string {
	// Fetch memory
	Mem, err := mem.VirtualMemory()
	if err != nil {
		return "Couldnt fetch memory info..."
	}
	usedMem := humanize.IBytes(Mem.Used)
	return usedMem
}

func GETMemoryTotal() string {
	// Fetch memory
	Mem, err := mem.VirtualMemory()
	if err != nil {
		return "Couldnt fetch memory info..."
	}
	totalMem := humanize.IBytes(Mem.Total)
	return totalMem
}
