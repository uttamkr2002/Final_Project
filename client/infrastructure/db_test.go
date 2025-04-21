package infrastructure

import (
	models "client/modelsWithInterface"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert" // for making test assertion
	"github.com/stretchr/testify/mock" // for creating mocks
)

type MockRowScanner struct {
	mock.Mock // mock.Mock provides methods for setting up expectations and return values
}

func (m *MockRowScanner) Scan(dest ...interface{}) error {
	args := m.Called(dest...) // we are the calling the mock function with arguments (dest..) which is basically row adress

	if len(dest) > 0 {
		if idptr, ok := dest[0].(*int); ok {
			*idptr = 12 //  storing 12 to that adress,
		}
	}
	return args.Error(0)
}

type MockSqlservice struct {
	mock.Mock
}

func (m *MockSqlservice) Ping() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockSqlservice) Closeconn() {
	m.Called()
}

func (m *MockSqlservice) QueryRow(query string, args ...interface{}) RowScanner {
	mockargs := []interface{}{query}
	mockargs = append(mockargs, args...)

	result := m.Called(mockargs...)
	return result.Get(0).(RowScanner)
}

func TestInsertsMetrics(t *testing.T) {
	metrics := models.Payload{
		Disk: models.DiskMetrics{
			Total:          100,
			Used:           50,
			IopsInProgress: 5,
		},
		Memory: models.MemoryMetrics{
			SwapTotal:    200,
			SwapUsed:     100,
			VirtualTotal: 300,
			VirtualUsed:  150,
			Buffers:      10,
			Cached:       20,
		},
		OS: models.OSMetrics{
			Uptime:          3600,
			Platform:        "Linux",
			PlatformVersion: "5.10.0",
		},
		CPU: models.CPUUsage{
			CPUUsage: 50.5,
		},
	}

	t.Run("Successful Insertion", func(t *testing.T) {
		mockDB := new(MockSqlservice)
		mockScanner := new(MockRowScanner)

		mockDB.On("Ping").Return(nil) // after calling Ping it should return nil.

		// Expected SQL query with RETURNING id
		mockDB.On("QueryRow", mock.Anything, mock.Anything, mock.Anything, mock.Anything, // mockDB.on , used to call QueryRow
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
			mock.Anything).Return(mockScanner)

		mockScanner.On("Scan", mock.Anything).Return(nil) // Above we are initializing 12

		id, err := InsertMetrics(mockDB, metrics)
		assert.NoError(t, err, "Expected no error during successful insert")
		assert.Equal(t, 12, id, "Expected returned ID to be 12") // wherther exist at that id or not is checked

		mockDB.AssertExpectations(t)
		mockScanner.AssertExpectations(t)
	})

	// assert.NoError := If no error then returns true.

	t.Run("Database Connection Error", func(t *testing.T) {
		mockDB := new(MockSqlservice)
		mockDB.On("Ping").Return(errors.New("Failed to connect to DB"))

		id, err := InsertMetrics(mockDB, metrics)
		assert.Error(t, err, "Expected an error when database connection fails")
		assert.Equal(t, 0, id, "Expected ID to be 0 when connection fails")
		assert.Contains(t, err.Error(), "database connection error", "Expected specific error message")

		mockDB.AssertExpectations(t)
	})

	t.Run("Error in Inserting Data", func(t *testing.T) {
		mockDB := new(MockSqlservice)
		mockScanner := new(MockRowScanner)

		mockDB.On("Ping").Return(nil)
		mockDB.On("QueryRow", mock.Anything, mock.Anything, mock.Anything, mock.Anything,
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
			mock.Anything).Return(mockScanner)

		mockScanner.On("Scan", mock.Anything).Return(errors.New("Insert Error"))

		id, err := InsertMetrics(mockDB, metrics)
		assert.Error(t, err, "Expected error when insert fails")
		assert.Equal(t, 0, id, "Expected ID to be 0 when insert fails")
		assert.Contains(t, err.Error(), "error inserting metrics", "Expected specific error message")

		mockDB.AssertExpectations(t)
		mockScanner.AssertExpectations(t)
	})
}
