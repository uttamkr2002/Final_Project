package modelsWithInterface

import (
	"fmt"
	"testing"
)

// table driven test
func TestMemoryCollectMetrics(t *testing.T) {
	MemoryMetric := &MemoryMetrics{}

	err := MemoryMetric.CollectMetrics()
	if err != nil {
		t.Errorf("We Encountered Error in %v", err)
		return
	}
}

var testCasesMemory = []struct {
	id       int
	expected MemoryMetrics
	actual   MemoryMetrics
	Result   bool
}{
	{1, MemoryMetrics{4096, 1024, 8192, 4096, 512, 1024}, MemoryMetrics{4096, 1024, 8192, 4096, 512, 1024}, true},  //  Should pass
	{2, MemoryMetrics{4096, 1024, 8192, 4096, 512, 1024}, MemoryMetrics{4096, 1024, 8192, 4000, 512, 1024}, false}, //  VirtualUsed mismatch
	{3, MemoryMetrics{0, 0, 0, 0, 0, 0}, MemoryMetrics{0, 0, 0, 0, 0, 0}, true},                                    // ✅ Should pass (zero values)
	{4, MemoryMetrics{4096, 1024, 8192, 4096, 512, 1024}, MemoryMetrics{4000, 1024, 8192, 4096, 512, 1024}, false}, //  SwapTotal mismatch
	{5, MemoryMetrics{4096, 1024, 8192, 4096, 512, 1024}, MemoryMetrics{4096, 1024, 8192, 4096, 500, 1024}, false}, //  Buffers mismatch
}

func TestMemory(t *testing.T) {
	for _, tc := range testCasesMemory {
		t.Run(fmt.Sprintf("Test-%d", tc.id), func(t *testing.T) {
			res := tc.expected == tc.actual
			if res != tc.Result {
				t.Errorf("❌ Test ID: %d failed: Expected %+v, but got %+v", tc.id, tc.expected, tc.actual)
			} else {
				fmt.Printf("✅ Test ID: %d passed\n", tc.id)
			}
		})
	}
}
