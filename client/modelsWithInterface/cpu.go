package modelsWithInterface

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
)

type CPUUsage struct {
	CPUUsage float64 `json:"cpu_usage"`
}

// GetCPUUsage fetches the CPU usage using the gopsutil/cpu package
func GetCPUUsage() (CPUUsage, error) {
	// Fetch the CPU usage percentage
	cpuPercent, err := cpu.Percent(0, false)
	fmt.Println("Hi, uTAM, CPUpERcent", cpuPercent)
	if err != nil {
		return CPUUsage{}, err
	}
	fmt.Println("Line 132", cpuPercent)
	// Return the CPU usage
	return CPUUsage{
		CPUUsage: cpuPercent[0],
	}, nil
}

// implement the method
func (cpu1 *CPUUsage) CollectMetrics() error {
	temp, err := GetCPUUsage() // returning struct and error
	if err != nil {
		// we can do panic, if we want to stop the execution here only
		// we can sue log , if we want to debug with time stamp
		// otherwise we can use normal println
		fmt.Println(err)
		return err
	}

	cpu1.CPUUsage = temp.CPUUsage //
	fmt.Println("Line 39, cpu.go", cpu1)
	return nil
}
