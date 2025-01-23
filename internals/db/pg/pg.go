package pg


import (
	"context"
	"fmt"
	"log"
	"time"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/jackc/pgx/v5/pgxpool"
)



// Pool is the global database connection pool
var Pool *pgxpool.Pool
var GDB *gorm.DB
func ConnectDatabase() {
	var err error

	dbHost := viper.GetString("pgdatabase.host")
	dbPort := viper.GetInt("pgdatabase.port")
	dbUser := viper.GetString("pgdatabase.user")
	dbPassword := viper.GetString("pgdatabase.password")
	dbName := viper.GetString("pgdatabase.dbname")
	dbSSLMode := viper.GetString("pgdatabase.sslmode")

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode)

	GDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connection established")
}

// NewDB function to create a new DB connection
// func InitializePostgresDB() ( error) {
//     connStr := "host=localhost port=5432 user=postgres password=mysecretpassword dbname=postgres sslmode=disable"
//     // Open the connection
//     DB, err := sql.Open("postgres", connStr)
//     if err != nil {
//         return fmt.Errorf("error opening database: %w", err)
//     }

//     // Check if the connection is alive
//     err = DB.Ping()
//     if err != nil {
//         return fmt.Errorf("error pinging database: %w", err)
//     }

//     fmt.Println("Successfully connected to the database!")

//     return nil
// }





// Initialize sets up the connection pool
func Initialize() {
	dbHost := viper.GetString("pgdatabase.host")
	dbPort := viper.GetString("pgdatabase.port")
	dbUser := viper.GetString("pgdatabase.user")
	dbPassword := viper.GetString("pgdatabase.password")
	dbName := viper.GetString("pgdatabase.dbname")
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	Pool, err = pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("Failed to initialize connection pool: %v", err)
	}

	// Test the connection
	if err := Pool.Ping(ctx); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}
	log.Println("Database connection pool initialized successfully.")
}
