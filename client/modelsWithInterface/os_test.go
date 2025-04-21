//package modelsWithInterface

//import "testing"

// implementing fuzz test things to keep in mind
// --> function name start with FuzzFunctionName
// --> The parameter in the function should contain (f *testing.F)
// --> No Return Type
// --> It continuously manipulates the Input to find Bugs
// --> Initialy we need to give some seed corpus in the form of f.Add()// after running this It continuously manipulates the input and test for the
// different types of edge case and report failure to the user
// to run the code we need to write go test -fuzz=functionName
package modelsWithInterface

import (
	"fmt"
	"testing"
)

// Table-Driven Testing for OSMetrics
var testcasesOS = []struct {
	id       int
	expected OSMetrics
	actual   OSMetrics
}{
	{
		id:       1,
		expected: OSMetrics{Uptime: 100, Platform: "Windows", PlatformVersion: "10.0"},
		actual:   OSMetrics{Uptime: 100, Platform: "Windows", PlatformVersion: "10.0"}, // Should pass
	},
	{
		id:       2,
		expected: OSMetrics{Uptime: 50000, Platform: "Linux", PlatformVersion: "5.4.0"},
		actual:   OSMetrics{Uptime: 50000, Platform: "Linux", PlatformVersion: "5.10.0"}, //  Should fail (Version mismatch)
	},
	{
		id:       3,
		expected: OSMetrics{},
		actual:   OSMetrics{}, //  Should pass (Zero values)
	},
	{
		id:       4,
		expected: OSMetrics{Uptime: 200, Platform: "macOS", PlatformVersion: "11.3"},
		actual:   OSMetrics{Uptime: 200, Platform: "macOS", PlatformVersion: "12.0"}, //  Should fail (Version mismatch)
	},
}

// Table-driven Unit Test for OSMetrics
func TestOSMetrics(t *testing.T) {
	for _, tc := range testcasesOS {
		t.Run(fmt.Sprintf("Test-%d", tc.id), func(t *testing.T) {
			if tc.expected != tc.actual {
				t.Errorf("❌ Test ID: %d failed: Expected %+v, but got %+v", tc.id, tc.expected, tc.actual)
			} else {
				fmt.Printf("✅ Test ID: %d passed\n", tc.id)
			}
		})
	}
}

// Function to test (replace with actual logic if needed)
func ProcessOSMetrics(metrics OSMetrics) string {
	// Example logic: Just return a formatted string
	return fmt.Sprintf("OS: %s %s, Uptime: %d", metrics.Platform, metrics.PlatformVersion, metrics.Uptime)
}

// Fuzz Test for OSMetrics
func FuzzOSMetrics(f *testing.F) {
	// Seed corpus (initial test values)
	f.Add(uint64(100), "Windows", "10.0")
	f.Add(uint64(200000), "Linux", "5.4.0")
	f.Add(uint64(0), "macOS", "11.3")

	// Fuzzing function
	f.Fuzz(func(t *testing.T, uptime uint64, platform string, version string) {
		// Create an OSMetrics instance with fuzzed values
		metrics := OSMetrics{
			Uptime:          uptime,
			Platform:        platform,
			PlatformVersion: version,
		}

		// Print the fuzzed test values
		fmt.Printf("🔍 Fuzz Testing with input: %+v\n", metrics)

		// Process the OS Metrics
		result := ProcessOSMetrics(metrics)

		// Check Unexpected test
		if result == "" {
			t.Errorf("❌ Unexpected empty result for input: %+v", metrics)
		}
	})
}
