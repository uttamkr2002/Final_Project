package models

import (

	// "github.com/shirou/gopsutil/cpu"

	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

// After structuring we are able to Populate.
type Payload struct {
	Disk   DiskMetrics   `json:"disk"`
	Memory MemoryMetrics `json:"memory"`
	OS     OSMetrics     `json:"OS"`
	CPU    CPUUsage      `"json:CPU"`
}

type DiskMetrics struct {
	Total          uint64 `json:"total"`
	Used           uint64 `json:"used"`
	IopsInProgress uint64 `json:"iopsInProgress"`
}

type MemoryMetrics struct {
	SwapTotal    uint64 `json:"swap_total"`
	SwapUsed     uint64 `json:"swap_used"`
	VirtualTotal uint64 `json:"virtual_total"`
	VirtualUsed  uint64 `json:"virtual_used"`
	Buffers      uint64 `json:"buffers"`
	Cached       uint64 `json:"cached"`
}

type OSMetrics struct {
	Uptime          uint64 `json:"uptime"`
	Platform        string `json:"platform"`
	PlatformVersion string `json:"platformVersion"`
}

type CPUUsage struct {
	CPUUsage float64 `json:"cpu_usage"`
}

func CollectMetrics() (Payload, error) {
	diskMetrics, err := getDiskMetrics()
	if err != nil {
		return Payload{}, err
	}

	memoryMetrics, err := getMemoryMetrics()
	if err != nil {
		return Payload{}, err
	}

	osMetrics, err := getOSMetrics()
	if err != nil {
		return Payload{}, err
	}

	cpuUsage, err := GetCPUUsage()
	if err != nil {
		return Payload{}, err
	}
	return Payload{
		Disk:   diskMetrics,
		Memory: memoryMetrics,
		OS:     osMetrics,
		CPU:    cpuUsage,
	}, nil
}

func getDiskMetrics() (DiskMetrics, error) {
	diskUsage, err := disk.Usage("/")
	if err != nil {
		return DiskMetrics{}, err
	}

	iopsInProgress, err := disk.IOCounters()
	if err != nil {
		return DiskMetrics{}, err
	}

	var totalIops uint64
	for _, iops := range iopsInProgress {
		totalIops += iops.IopsInProgress
	}

	return DiskMetrics{
		Total:          diskUsage.Total,
		Used:           diskUsage.Used,
		IopsInProgress: totalIops,
	}, nil
}

func getMemoryMetrics() (MemoryMetrics, error) {
	memoryStats, err := mem.VirtualMemory()
	if err != nil {
		return MemoryMetrics{}, err
	}

	swapStats, err := mem.SwapMemory()
	if err != nil {
		return MemoryMetrics{}, err
	}

	return MemoryMetrics{
		SwapTotal:    swapStats.Total,
		SwapUsed:     swapStats.Used,
		VirtualTotal: memoryStats.Total,
		VirtualUsed:  memoryStats.Used,
		Buffers:      memoryStats.Buffers,
		Cached:       memoryStats.Cached,
	}, nil
}

func getOSMetrics() (OSMetrics, error) {
	hostInfo, err := host.Info()
	if err != nil {
		return OSMetrics{}, err
	}

	return OSMetrics{
		Uptime:          hostInfo.Uptime,
		Platform:        hostInfo.Platform,
		PlatformVersion: hostInfo.PlatformVersion,
	}, nil
}

// GetCPUUsage fetches the CPU usage using the gopsutil/cpu package
func GetCPUUsage() (CPUUsage, error) {
	// Fetch the CPU usage percentage
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		return CPUUsage{}, err
	}
	fmt.Println("Line 132", cpuPercent)
	// Return the CPU usage
	return CPUUsage{
		CPUUsage: cpuPercent[0],
	}, nil
}

// Difficult to test: Since the code directly calls system APIs, writing unit tests becomes harder.
// No flexibility: If you want to use another library (or mock metrics for testing), you need to modify the core functions.


// Interface are Used for code reusablity
// So here I can find that Collect metrics is common for collecting all the metrics