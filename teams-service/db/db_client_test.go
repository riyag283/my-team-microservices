package db

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/mock"
)

// Define a mock implementation of the DBClientInterface
type MockDBClientInterface struct {
    mock.Mock
}

func (m *MockDBClientInterface) Ping() error {
    args := m.Called()
    return args.Error(0)
}

func (m *MockDBClientInterface) Close() error {
    args := m.Called()
    return args.Error(0)
}

func (m *MockDBClientInterface) Exec(query string, args ...interface{}) (sql.Result, error) {
    argsList := m.Called(query, args)
    return argsList.Get(0).(sql.Result), argsList.Error(1)
}

func (m *MockDBClientInterface) Get(dest interface{}, query string, args ...interface{}) error {
    argsList := m.Called(dest, query, args)
    return argsList.Error(0)
}

func (m *MockDBClientInterface) Select(dest interface{}, query string, args ...interface{}) error {
    argsList := m.Called(dest, query, args)
    return argsList.Error(0)
}

func (m *MockDBClientInterface) QueryRow(query string, args ...interface{}) *sql.Row {
    argsList := m.Called(query, args)
    return argsList.Get(0).(*sql.Row)
}

func (m *MockDBClientInterface) QueryRowx(query string, args ...interface{}) *sqlx.Row {
    argsList := m.Called(query, args)
    return argsList.Get(0).(*sqlx.Row)
}

func (m *MockDBClientInterface) Query(query string, args ...interface{}) (*sql.Rows, error) {
    argsList := m.Called(query, args)
    return argsList.Get(0).(*sql.Rows), argsList.Error(1)
}

func TestInitialiseDBConnection(t *testing.T) {
    // Create a mock DBClientInterface instance
    mockDBClient := new(MockDBClientInterface)

    // Define the expected behavior of the Ping() method
    mockDBClient.On("Ping").Return(nil)

    // Define the expected behavior of the Close() method
    mockDBClient.On("Close").Return(nil)

    // Define the expected behavior of the Query() method
    mockDBClient.On("Query", mock.Anything, mock.Anything).Return(nil, errors.New("mock error"))

    // Inject the mock DBClientInterface instance into the DBClientInstance variable
    DBClientInstance = mockDBClient

    // Call the InitialiseDBConnection() function
    InitialiseDBConnection()

    // Assert that Ping() was called on the mock DBClientInterface instance
    mockDBClient.AssertCalled(t, "Ping")

    // Assert that Close() was called on the mock DBClientInterface instance
    mockDBClient.AssertCalled(t, "Close")

    // Assert that Query() was not called on the mock DBClientInterface instance
    mockDBClient.AssertNotCalled(t, "Query")
}
