package modelsWithInterface

import (
	"fmt"
	"testing"
)

// func TestDiskCollectionMetrics(t *testing.T) {
//    diskmetric := &DiskMetrics{}
//    err := diskmetric.CollectMetrics()
//    if err != nil{
// 	t.Errorf("We Encountered Error%v",err)
// 	return
//    }
// }

// Table driven unit Test
var testcasesDisk = []struct {
	id       int64
	expected DiskMetrics
	actual   DiskMetrics
	result   bool
}{
	{
		id:       1,
		expected: DiskMetrics{Total: 1000, Used: 500, IopsInProgress: 10},
		actual:   DiskMetrics{Total: 1000, Used: 500, IopsInProgress: 10}, // Should pass
		result:   true,
	},
	{
		id:       2,
		expected: DiskMetrics{Total: 1000, Used: 500, IopsInProgress: 10},
		actual:   DiskMetrics{Total: 1000, Used: 500, IopsInProgress: 0}, // Should fail
		result:   false,
	},
	{
		id:       3,
		expected: DiskMetrics{},
		actual:   DiskMetrics{}, //pass
		result:   true,
	},
	{
		id:       4,
		expected: DiskMetrics{Total: 1000, Used: 500, IopsInProgress: 10},
		actual:   DiskMetrics{Total: 1000, Used: 0, IopsInProgress: 10}, // Should fail
		result:   false,
	},
}

func TestDisk(t *testing.T) {
	for _, tc := range testcasesDisk {
		t.Run(fmt.Sprintf("Test-%d", tc.id), func(t *testing.T) {
			res := tc.expected == tc.actual
			if res != tc.result {
				t.Errorf("❌ Test ID: %d failed: Expected %+v, but got %+v", tc.id, tc.expected, tc.actual)
			} else {
				t.Logf("✅ Test ID: %d passed\n", tc.id)
			}
		})
	}
}
