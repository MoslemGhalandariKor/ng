package ora

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/godror/godror"
)

// var DB *sql.DB

// // Init initializes the connection to the Oracle database
// func Init() error {
// 	var err error
// 	dsn := "sadmin/mnzxa123@localhost:1521/nextgendb"
// 	DB, err = sql.Open("godror", dsn)
// 	if err != nil {
// 		return fmt.Errorf("error opening database: %w", err)
// 	}

// 	// Test the connection
// 	ctx := context.Background()
// 	if err = DB.PingContext(ctx); err != nil {
// 		log.Fatalf("Failed to connect to oracle database: %v", err)
// 	}

// 	log.Println("Successfully connected to Oracle database!")
// 	return nil
// }

// // QueryExample demonstrates a sample query
// func QueryExample() error {
// 	ctx := context.Background()
// 	rows, err := DB.QueryContext(ctx, "SELECT * FROM employees WHERE ROWNUM <= 5")
// 	if err != nil {
// 		return fmt.Errorf("error executing query: %w", err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var empID int
// 		var empName string
// 		if err := rows.Scan(&empID, &empName); err != nil {
// 			return fmt.Errorf("error scanning row: %w", err)
// 		}
// 		fmt.Printf("Employee ID: %d, Name: %s\n", empID, empName)
// 	}

// 	return nil
// }



type OraConfig struct {
    OracleUser                string `mapstructure:"oracle_user"`
    OraclePassword            string `mapstructure:"oracle_password"`
    OracleConnectString       string `mapstructure:"oracle_connect_string"`
	OracleInstantClientPath   string `mapstructure:"oracle_instant_client_path"`
    OracleMaxOpenConns        int    `mapstructure:"oracle_max_open_conns"`
    OracleMaxIdleConns        int    `mapstructure:"oracle_max_idle_conns"`
}
var OraDB *sql.DB

func Initialize(dsn string, maxOpenConns, maxIdleConns int) error {
	var err error

	// Open the connection using database/sql and godror.
	OraDB, err = sql.Open("godror", dsn)
	if err != nil {
		return fmt.Errorf("failed to open Oracle database connection: %w", err)
	}

	// Configure the connection pool settings.
	OraDB.SetMaxOpenConns(maxOpenConns)
	OraDB.SetMaxIdleConns(maxIdleConns)
	OraDB.SetConnMaxLifetime(0)

	// Verify that the connection is established.
	if err = OraDB.Ping(); err != nil {
		return fmt.Errorf("failed to connect to the Oracle database: %w", err)
	}

	fmt.Println("Connected to the Oracle database successfully!")
	return nil
}

func LogConnectionPoolStats() {
    stats := OraDB.Stats()
    log.Printf("Open connections: %d, Idle connections: %d", stats.OpenConnections, stats.Idle)
}

func CallProcedure(){
	ctx := context.Background()

	var p_response_code int
	var p_response_desc string

	q := `BEGIN 
		    PRC_TEST(:1, :2, :3, :4); 
		  END;`


	_, err := OraDB.ExecContext(ctx,
					           q,
						       "param",
							   "param",
							   sql.Out{Dest: &p_response_code},
							   sql.Out{Dest: &p_response_desc},
								)
								
    if err !=nil{
		log.Fatalf("Failed to execute procedure: %v", err)
	}
	fmt.Printf("Mapped Output P_RESPONSE_CODE: %d\n", p_response_code)
	fmt.Printf("Mapped Output P_RESPONSE_DESC: %s\n", p_response_desc)

}