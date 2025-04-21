package modelsWithInterface

import (
	"fmt"
	"testing"
)
// right now I have created only 1 testcase, if its good then I will update all the metrics code
var testCases = []struct {
	id       int64
	expected error
	res      bool
}{
	{1, nil, true}, // pass 
	//{1,err,false}, this situation is for error but we cannot predict that systemcall will fail
}
// It Test will only If System call fails
// nil values check,
func TestGetCPUUsage(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Test-%d", tc.id), func(t *testing.T) { 			
			_, err := GetCPUUsage()	// there is no use of the data that is returned by GetCPUUsages so we are using _		
			if tc.res == true { // 
				if err != tc.expected {
					t.Errorf("got Error in %v, Expected %v, Got %v", tc.id, tc.expected, err)
				}
			}

		})
	}
}

func TestCpuMetric(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Test-%d", tc.id), func(t *testing.T) {
			cpuObj, _ := GetCPUUsage()
			err1 := cpuObj.CollectMetrics()	

			if tc.res == true {
				if err1 != tc.expected {
					t.Errorf("got Error in %v, Expected %v, Got %v", tc.id, tc.expected, err1)
				}
			}

		})
	}
}
