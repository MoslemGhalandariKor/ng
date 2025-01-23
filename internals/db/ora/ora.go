package ora

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/godror/godror"
)

var DB *sql.DB

// Init initializes the connection to the Oracle database
func Init() error {
	var err error
	dsn := "sadmin/mnzxa123@localhost:1521/nextgendb"
	DB, err = sql.Open("godror", dsn)
	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}

	// Test the connection
	ctx := context.Background()
	if err = DB.PingContext(ctx); err != nil {
		log.Fatalf("Failed to connect to oracle database: %v", err)
	}

	log.Println("Successfully connected to Oracle database!")
	return nil
}

// QueryExample demonstrates a sample query
func QueryExample() error {
	ctx := context.Background()
	rows, err := DB.QueryContext(ctx, "SELECT * FROM employees WHERE ROWNUM <= 5")
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var empID int
		var empName string
		if err := rows.Scan(&empID, &empName); err != nil {
			return fmt.Errorf("error scanning row: %w", err)
		}
		fmt.Printf("Employee ID: %d, Name: %s\n", empID, empName)
	}

	return nil
}
