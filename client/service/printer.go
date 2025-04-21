package service

import (
	models "client/modelsWithInterface"
	"fmt"
)

func PrintMetrics(payload models.Payload) {
	fmt.Println("Disk Metrics:")
	fmt.Printf("Total: %d\nUsed: %d\nIOPS In Progress: %d\n\n", payload.Disk.Total, payload.Disk.Used, payload.Disk.IopsInProgress)

	fmt.Println("Memory Metrics:")
	fmt.Printf("Swap Total: %d\nSwap Used: %d\nVirtual Total: %d\nVirtual Used: %d\nBuffers: %d\nCached: %d\n\n",
		payload.Memory.SwapTotal, payload.Memory.SwapUsed, payload.Memory.VirtualTotal, payload.Memory.VirtualUsed,
		payload.Memory.Buffers, payload.Memory.Cached)

	fmt.Println("OS Metrics:")
	fmt.Printf("Uptime: %d\nPlatform: %s\nPlatform Version: %s\n",
		payload.OS.Uptime, payload.OS.Platform, payload.OS.PlatformVersion)

	fmt.Println("CPU Metrics:")
	fmt.Printf("CPU Percentage %f\n",
		payload.CPU.CPUUsage)
}
