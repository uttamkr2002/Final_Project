package infrastructure

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// RowScanner interface for scanning rows
type RowScanner interface {
	Scan(dest ...interface{}) error // slice of interface taking input adresss
}

// SqlRow wraps sql.Row to implement RowScanner
type SqlRow struct {
	row *sql.Row
}

// If I try to return the struct, then I will face difficulty while mocking

// Scan implements the RowScanner interface
func (r *SqlRow) Scan(dest ...interface{}) error {
	return r.row.Scan(dest...)
}

// SqlServices defines the interface for SQL operations
type SqlServices interface {
	Ping() error
	Closeconn()
	QueryRow(query string, args ...interface{}) RowScanner
}

// SqlClient represents the SQL client type
type SqlClient struct {
	db *sql.DB
}

// Ping checks the database connection
func (dbClient *SqlClient) Ping() error {
	if dbClient.db == nil {
		return fmt.Errorf("database connection is not initialized")
	}
	return dbClient.db.Ping()
}

// Closeconn closes the database connection
func (dbClient *SqlClient) Closeconn() {
	if dbClient.db != nil {
		dbClient.db.Close()
		fmt.Println("✅ Database connection closed")
	}
}

// QueryRow executes a query that returns a single row and implements RowScanner
func (dbClient *SqlClient) QueryRow(query string, args ...interface{}) RowScanner {
	if dbClient.db == nil {
		// Return a SqlRow with a nil row, which will fail on Scan
		return &SqlRow{row: nil}
	}

	row := dbClient.db.QueryRow(query, args...)
	return &SqlRow{row: row}
}

func Loaddotenv() error {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("⚠️ Warning: Error loading .env file:", err)
		return err
	}
	fmt.Println("✅ .env file loaded successfully")
	return nil
}

func ReturnConnectionString(user, pass, dbName, sslMode string) string {
	fmt.Println(user, pass, dbName, sslMode)

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", user, pass, dbName, sslMode)
	fmt.Println("🔗 Connection String:", connectionString)
	return connectionString
}

// InitDb initializes the database connection
func InitDb() (*SqlClient, error) {
	fmt.Println("🔄 Initializing database connection...")

	// loading dotenv
	Loaddotenv()

	user := os.Getenv("USER")
	pass := os.Getenv("Password")
	dbName := os.Getenv("Dbname")
	sslMode := os.Getenv("Sslmode")
	fmt.Println(user, pass, dbName, sslMode)

	connectionString := ReturnConnectionString(user, pass, dbName, sslMode)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Printf("❌ Error opening DB connection: %v\n", err)
		return nil, err
	}

	dbClient := &SqlClient{db: db}

	if err = dbClient.Ping(); err != nil {
		fmt.Printf("❌ Error pinging DB: %v\n", err)
		db.Close()
		return nil, err
	}

	fmt.Println("✅ Successfully connected to the database!")
	return dbClient, nil
}
