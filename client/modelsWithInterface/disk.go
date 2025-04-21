package modelsWithInterface

import (
	"fmt"

	"github.com/shirou/gopsutil/disk"
)

type DiskMetrics struct {
	Total          uint64 `json:"total"`
	Used           uint64 `json:"used"`
	IopsInProgress uint64 `json:"iopsInProgress"`
}

// implement the common method

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
	fmt.Println(diskUsage)
	return DiskMetrics{
		Total:          diskUsage.Total,
		Used:           diskUsage.Used,
		IopsInProgress: totalIops,
	}, nil

}

func (disk1 *DiskMetrics) CollectMetrics() error{
	temp, err := getDiskMetrics()
	//fmt.Println(temp)
	
	if err != nil {
		fmt.Println("We Encountered an Error")
		return err
	}
	//disk1 = &temp // not works ?
	disk1.Total = temp.Total // this works fine
	disk1.Used = temp.Used
	disk1.IopsInProgress = temp.IopsInProgress
	fmt.Println(disk1)
	return nil
}
