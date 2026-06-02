package src

import (
	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/disk"
)

func GetDiskTotal() string {
	Disk, err := disk.Usage("/")
	if err != nil {
		return "Error occured while finding total disk space"
	}
	TotalSpace := humanize.IBytes(Disk.Total)
	return TotalSpace
}

func GetDiskUsed() string {
	Disk, err := disk.Usage("/")
	if err != nil {
		return "Error occured while finding total disk space"
	}
	TotalSpace := humanize.IBytes(Disk.Used)
	return TotalSpace
}
