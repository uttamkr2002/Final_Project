package modelsWithInterface

import (
	"errors"
	"fmt"
	"testing"
)

// Function to process payload and return an error if invalid
func ProcessPayload(p Payload) error {
	if (p.Disk == DiskMetrics{}) { // Checking empty struct instead of nil
		return errors.New("❌ Error: Disk metrics are missing")
	}
	if (p.Memory == MemoryMetrics{}) { // Checking empty struct instead of nil
		return errors.New("❌ Error: Memory metrics are missing")
	}
	if p.OS.Platform == "" || p.OS.PlatformVersion == "" {
		return errors.New("❌ Error: OS metrics are incomplete")
	}
	if p.CPU.CPUUsage < 0 {
		return errors.New("❌ Error: CPU usage cannot be negative")
	}
	return nil
}

// Table-driven test case for ProcessPayload
func TestProcessPayload(t *testing.T) {
	testCases := []struct {
		name    string
		payload Payload
		wantErr bool
	}{
		{
			name: "✅ Valid Payload",
			payload: Payload{
				Disk:   DiskMetrics{Total: 1000, Used: 500, IopsInProgress: 10},
				Memory: MemoryMetrics{SwapTotal: 4096, SwapUsed: 1024, VirtualTotal: 8192, VirtualUsed: 4096, Buffers: 512, Cached: 1024},
				OS:     OSMetrics{Platform: "Linux", PlatformVersion: "5.4", Uptime: 50000},
				CPU:    CPUUsage{CPUUsage: 25.5},
			},
			wantErr: false,
		},
		{
			name: "❌ Missing Disk Metrics",
			payload: Payload{
				Disk:   DiskMetrics{}, // Empty struct instead of nil
				Memory: MemoryMetrics{SwapTotal: 4096, SwapUsed: 1024, VirtualTotal: 8192, VirtualUsed: 4096, Buffers: 512, Cached: 1024},
				OS:     OSMetrics{Platform: "Linux", PlatformVersion: "5.4", Uptime: 50000},
				CPU:    CPUUsage{CPUUsage: 25.5},
			},
			wantErr: true,
		},
		{
			name: "❌ Missing Memory Metrics",
			payload: Payload{
				Disk:   DiskMetrics{Total: 1000, Used: 500, IopsInProgress: 10},
				Memory: MemoryMetrics{}, // Empty struct instead of nil
				OS:     OSMetrics{Platform: "Linux", PlatformVersion: "5.4", Uptime: 50000},
				CPU:    CPUUsage{CPUUsage: 25.5},
			},
			wantErr: true,
		},
		{
			name: "❌ Incomplete OS Metrics",
			payload: Payload{
				Disk:   DiskMetrics{Total: 1000, Used: 500, IopsInProgress: 10},
				Memory: MemoryMetrics{SwapTotal: 4096, SwapUsed: 1024, VirtualTotal: 8192, VirtualUsed: 4096, Buffers: 512, Cached: 1024},
				OS:     OSMetrics{Platform: "", PlatformVersion: "", Uptime: 50000}, // Empty OS fields
				CPU:    CPUUsage{CPUUsage: 25.5},
			},
			wantErr: true,
		},
		{
			name: "❌ Negative CPU Usage",
			payload: Payload{
				Disk:   DiskMetrics{Total: 1000, Used: 500, IopsInProgress: 10},
				Memory: MemoryMetrics{SwapTotal: 4096, SwapUsed: 1024, VirtualTotal: 8192, VirtualUsed: 4096, Buffers: 512, Cached: 1024},
				OS:     OSMetrics{Platform: "Linux", PlatformVersion: "5.4", Uptime: 50000},
				CPU:    CPUUsage{CPUUsage: -10}, // Invalid negative CPU usage
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := ProcessPayload(tc.payload)
			if (err != nil) != tc.wantErr {
				t.Errorf("Test %q failed: expected error: %v, got error: %v", tc.name, tc.wantErr, err)
			} else if err != nil {
				fmt.Println(err) // Print error for debugging
			} else {
				fmt.Println("✅ Test passed:", tc.name)
			}
		})
	}
}
