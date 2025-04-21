// this file contains the structure that I am returning
// After structuring we are able to Populate.

package modelsWithInterface

import "fmt"

type Payload struct {
	Disk   DiskMetrics   `json:"disk"`   // if we put here *DiskMetrics?, then below we will give &DiskMetrics{}
	Memory MemoryMetrics `json:"memory"` // is this correct or we need to give pointer here? confused where to
	OS     OSMetrics     `json:"OS"`     // give pointer and where to not?
	CPU    CPUUsage      `json:"CPU"`
}

func (P1 *Payload) CollectMetricsforPayload() {
	// link btw interface and the struct that are implementing the interface
	// var m1 MetricsCollection  // m1 can hold any type of struct(cpu,os,disk,..) that implements MetricsColelction interface.

	var m1 MetricsCollection // m1 is variable of type interface than can hold any data type cpu, disk,os,memory
	// Collect Disk Metrics
	m1 = &P1.Disk
	err := m1.CollectMetrics()
	if err != nil {
		fmt.Println("Error while fetching Disk metrics:", err)
		return
	}
	//

	// Collect Memory Metrics
	m1 = &P1.Memory
	err = m1.CollectMetrics()
	if err != nil {
		fmt.Println("Error while fetching Memory metrics:", err)
		return
	}

	// Collect OS Metrics
	m1 = &P1.OS
	err = m1.CollectMetrics()
	if err != nil {
		fmt.Println("Error while fetching OS metrics:", err)
		return
	}

	// Collect CPU Metrics
	m1 = &P1.CPU
	err = m1.CollectMetrics()
	if err != nil {
		fmt.Println("Error while fetching CPU metrics:", err)
		return
	}

	// **Earlier we are using this what is wrong in this still getting confused, Is it because we want to implement polymorphism
	// P1.Disk.CollectMetrics()
	// P1.Memory.CollectMetrics()
	// P1.OS.CollectMetrics()
	//P1.CPU.CollectMetrics()
}

// no need of collect

// if we are using P1.Disk = DiskMetrics{} then its just giving empty struct
