package modelsWithInterface

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

type MemoryMetrics struct {
	SwapTotal    uint64 `json:"swap_total"`
	SwapUsed     uint64 `json:"swap_used"`
	VirtualTotal uint64 `json:"virtual_total"`
	VirtualUsed  uint64 `json:"virtual_used"`
	Buffers      uint64 `json:"buffers"`
	Cached       uint64 `json:"cached"`
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




// implement the common method

func (m1 *MemoryMetrics) CollectMetrics()error{
   temp , err := getMemoryMetrics()
   if err != nil{
	fmt.Println(err)
	return err
   }
   //m1 = &temp  // why it is not working? giving empty or default values
    m1.Buffers = temp.Buffers
	m1.Cached = temp.Cached
	m1.SwapTotal = temp.SwapTotal
	m1.SwapUsed = temp.SwapUsed
	m1.VirtualTotal = temp.VirtualTotal
	m1.VirtualUsed = temp.VirtualUsed
	return nil
}  
