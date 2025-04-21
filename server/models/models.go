package models

// "github.com/shirou/gopsutil/cpu"

// After structuring we are able to Populate.
type Payload struct {
	Disk   DiskMetrics   `json:"disk" bson:"disk"`
	Memory MemoryMetrics `json:"memory"`
	OS     OSMetrics     `json:"OS"`
	CPU    CPUUsage      `"json:CPU"`
}

type DiskMetrics struct {
	Total          uint64 `json:"total" bson:"total"`
	Used           uint64 `json:"used" bson:"used"`
	IopsInProgress uint64 `json:"iopsInProgress" bson:"iopsInProgress"`
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
