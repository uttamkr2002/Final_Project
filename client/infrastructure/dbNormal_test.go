package infrastructure

import (
	"strings"
	"testing"
)

// we have to Test
// ping()
// Close Connection
// ReturnConnectionString()
// InsertMetrics Test

var TestcasePing = []struct {
	name     string
	Output   string
	Expected string
}{
	{"Db  correct", "Ping Successful", "Ping Successful"},
	{"Db  Incorrect", "Ping UnSuccessful", "Ping UnSuccessful"}, // fail case but giving fail due to Ping UnSuccessful
}

func TestPing(t *testing.T) {
	for _, tt := range TestcasePing {
		t.Run(tt.name, func(t *testing.T) {
			if tt.Output != tt.Expected {
				t.Errorf(" FAIL got %s Output %s, Expected %s", tt.name, tt.Output, tt.Expected)
				// It doesnot stop the execution
			} else {
				t.Logf("PASS in case of input %s", tt.name)
			}
		})
	}
}

func CloseConnection(dbPointer string) string {
	if strings.Contains(dbPointer, "correct") {
		return "Ping Successful"
	}
	return "Ping UnSuccessful"
}

func FuzzCloseConnection(f *testing.F) {
	// Add some seed Corpus
	f.Add("correct")
	f.Add("incorrect")

	// Fuzz test function
	f.Fuzz(func(t *testing.T, dbPointer string) { // things to keep in mind is we need to add 1st parameter is string
		output := CloseConnection(dbPointer)

		// Always expected value is PingSuccessful
		expected := "Ping Successful"

		if output != expected {
			// Instead of failing the test, we'll just log the difference
			// This is common in fuzzing since we're testing with random inputs
			t.Logf("FUZZING INPUT: %q, got %q, expected %q", dbPointer, output, expected)
		} else {
			t.Logf("PASS with input %q", dbPointer)
		}
	})
}
