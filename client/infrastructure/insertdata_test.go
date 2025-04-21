package infrastructure

import (
	models "client/modelsWithInterface"
	"fmt"
	"log"
	"testing"
)

// Testcase
// 1.  All value in Range , Expected Error :None
// 2. Zero Can be inserted,  Expected Error : nO
// 3. My datatype is very large number, Expected Error : No
// 4. Negative Values, Expected Error : Yes
// 5. If One struct is Empty : Expected Error Yes

var tests = []struct {
	name        string
	payload     models.Payload
	expectedErr bool
}{
	{
		name: "Success - All valid parameters",
		payload: models.Payload{
			Disk: models.DiskMetrics{
				Total:          100,
				Used:           50,
				IopsInProgress: 10,
			},
			Memory: models.MemoryMetrics{
				SwapTotal:    1000,
				SwapUsed:     200,
				VirtualTotal: 2000,
				VirtualUsed:  1000,
				Buffers:      300,
				Cached:       400,
			},
			OS: models.OSMetrics{
				Uptime:          3600,
				Platform:        "Linux",
				PlatformVersion: "5.4.0",
			},
			CPU: models.CPUUsage{
				CPUUsage: 25.5,
			},
		},
		expectedErr: false,
	},
	{
		name: "Edge Case - Zero values",
		payload: models.Payload{
			Disk: models.DiskMetrics{
				Total:          0,
				Used:           0,
				IopsInProgress: 0,
			},
			Memory: models.MemoryMetrics{
				SwapTotal:    0,
				SwapUsed:     0,
				VirtualTotal: 0,
				VirtualUsed:  0,
				Buffers:      0,
				Cached:       0,
			},
			OS: models.OSMetrics{
				Uptime:          0,
				Platform:        "",
				PlatformVersion: "",
			},
			CPU: models.CPUUsage{
				CPUUsage: 0,
			},
		},
		expectedErr: false,
	},
	{
		name: "Edge Case - Null string values",
		payload: models.Payload{
			Disk: models.DiskMetrics{
				Total:          100,
				Used:           50,
				IopsInProgress: 10,
			},
			Memory: models.MemoryMetrics{
				SwapTotal:    1000,
				SwapUsed:     200,
				VirtualTotal: 2000,
				VirtualUsed:  1000,
				Buffers:      300,
				Cached:       400,
			},
			OS: models.OSMetrics{
				Uptime: 3600,
				// Platform and PlatformVersion are intentionally not set
			},
			CPU: models.CPUUsage{
				CPUUsage: 25.5,
			},
		},
		expectedErr: false,
	},
	{
		name: "Edge Case - Only required fields",
		payload: models.Payload{
			// Minimal required fields, others will be zero values
			Disk: models.DiskMetrics{
				Total: 100,
				Used:  50,
			},
			Memory: models.MemoryMetrics{
				VirtualTotal: 2000,
				VirtualUsed:  1000,
			},
			OS: models.OSMetrics{
				Uptime: 3600,
			},
		},
		expectedErr: false,
	},
	{
		name: "Extreme large values",
		payload: models.Payload{
			Disk: models.DiskMetrics{
				Total:          9999999999999,
				Used:           9999999999999,
				IopsInProgress: 9999999999999,
			},
			Memory: models.MemoryMetrics{
				SwapTotal:    9999999999999,
				SwapUsed:     9999999999999,
				VirtualTotal: 9999999999999,
				VirtualUsed:  9999999999999,
				Buffers:      9999999999999,
				Cached:       9999999999999,
			},
			OS: models.OSMetrics{
				Uptime:          9999999999999,
				Platform:        "This is an extremely long platform name that might exceed the database column size limit depending on how the schema is defined",
				PlatformVersion: "This is an extremely long platform version that might exceed the database column size limit depending on how the schema is defined",
			},
			CPU: models.CPUUsage{
				CPUUsage: 999999.99,
			},
		},
		expectedErr: false,
	},
	// {
	// 	name: "Negative values",
	// 	payload: models.Payload{
	// 		Disk: models.DiskMetrics{
	// 			Total:          -100,
	// 			Used:           -50,
	// 			IopsInProgress: -10,
	// 		},
	// 		Memory: models.MemoryMetrics{
	// 			SwapTotal:    -1000,
	// 			SwapUsed:     -200,
	// 			VirtualTotal: -2000,
	// 			VirtualUsed:  -1000,
	// 			Buffers:      -300,
	// 			Cached:       -400,
	// 		},
	// 		OS: models.OSMetrics{
	// 			Uptime:          -3600,
	// 			Platform:        "Linux",
	// 			PlatformVersion: "5.4.0",
	// 		},
	// 		CPU: models.CPUUsage{
	// 			CPUUsage: -25.5,
	// 		},
	// 	},
	// 	expectedErr: true,
	// },
	//Added new valid test case 1
	{
		name: "Success - Windows system",
		payload: models.Payload{
			Disk: models.DiskMetrics{
				Total:          500,
				Used:           250,
				IopsInProgress: 5,
			},
			Memory: models.MemoryMetrics{
				SwapTotal:    4000,
				SwapUsed:     1000,
				VirtualTotal: 8000,
				VirtualUsed:  3500,
				Buffers:      200,
				Cached:       800,
			},
			OS: models.OSMetrics{
				Uptime:          86400, // 1 day
				Platform:        "Windows",
				PlatformVersion: "10.0.19042",
			},
			CPU: models.CPUUsage{
				CPUUsage: 47.2,
			},
		},
		expectedErr: false,
	},
	// Added new valid test case 2
	{
		name: "Success - MacOS system",
		payload: models.Payload{
			Disk: models.DiskMetrics{
				Total:          1000,
				Used:           600,
				IopsInProgress: 8,
			},
			Memory: models.MemoryMetrics{
				SwapTotal:    2000,
				SwapUsed:     500,
				VirtualTotal: 16000,
				VirtualUsed:  12000,
				Buffers:      1000,
				Cached:       2000,
			},
			OS: models.OSMetrics{
				Uptime:          172800, // 2 days
				Platform:        "Darwin",
				PlatformVersion: "20.6.0",
			},
			CPU: models.CPUUsage{
				CPUUsage: 32.8,
			},
		},
		expectedErr: false,
	},
	// Added empty CPU struct case
	{
		name: "Edge Case - Empty CPU struct",
		payload: models.Payload{
			Disk: models.DiskMetrics{
				Total:          100,
				Used:           50,
				IopsInProgress: 10,
			},
			Memory: models.MemoryMetrics{
				SwapTotal:    1000,
				SwapUsed:     200,
				VirtualTotal: 2000,
				VirtualUsed:  1000,
				Buffers:      300,
				Cached:       400,
			},
			OS: models.OSMetrics{
				Uptime:          3600,
				Platform:        "Linux",
				PlatformVersion: "5.4.0",
			},
			CPU: models.CPUUsage{}, // Empty CPU struct
		},
		expectedErr: false,
	},
}

func TestInsertMetrics(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			dbClient, err := InitDb()
			if err != nil {
				log.Fatalf(" Error initializing database: %v", err)
				return
			}

			// Call the function being tested
			i, err := InsertMetrics(dbClient, tt.payload)
			fmt.Println(i)

			// Check if the error result matches what was expected
			if (err != nil) != tt.expectedErr {
				t.Errorf("InsertMetrics() error = %v, expectedErr %v", err, tt.expectedErr)
			}

			// If you need to check specific error types or messages:
			if tt.expectedErr && err == nil {
				t.Errorf("InsertMetrics() expected an error but got nil")
			} else if !tt.expectedErr && err != nil {
				t.Errorf("InsertMetrics() expected no error but got: %v", err)
			}

		})
	}
}
