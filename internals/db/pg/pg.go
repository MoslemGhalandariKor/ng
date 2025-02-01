package pg

import (
	"context"
	"fmt"
	"log"
	"time"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgConfig struct {
	PostgresHost              string `mapstructure:"postgres_host"`
	PostgresUser              string `mapstructure:"postgres_user"`
	PostgresPassword          string `mapstructure:"postgres_password"`
	PostgresSslMode           string `mapstructure:"postgres_sslmode"`
	PostgresDbName            string `mapstructure:"postgres_dbname"`
	PostgresPort              int    `mapstructure:"postgres_port"`
	PostgresMaxOpenConns      int32    `mapstructure:"oracle_max_open_conns"`
	PostgresMaxIdleConns      int32    `mapstructure:"oracle_max_idle_conns"`
	PostgresMaxConnIdleTime   int    `mapstructure:"oracle_max_idle_conns"`
	PostgresMaxConnLifeTime   int    `mapstructure:"oracle_max_idle_conns"`
	PostgresHealthCheckPeriod int    `mapstructure:"oracle_max_idle_conns"`
}

// DB holds the connection pool.
var PgDB *pgxpool.Pool

var GDB *gorm.DB

// Connect sets up a connection pool to the PostgreSQL database with custom configuration.
func Initialize(connString string, maxOpenConns int32, maxIdleConns int32,) error {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatalf("Unable to parse connection string: %v\n", err)
	}

	// Customize the pool settings here:
	config.MaxConns = maxOpenConns             // Maximum number of connections in the pool
	config.MinConns = maxIdleConns             // Minimum number of connections in the pool
	config.MaxConnIdleTime =   time.Minute * 5   // Maximum time a connection can be idle before being closed
	config.MaxConnLifetime =   time.Hour * 1     // Maximum time a connection can be reused before being closed
	config.HealthCheckPeriod = time.Minute * 1 // Period for health checks to remove unhealthy connections

	// Create a new connection pool with the configured settings
	PgDB, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}

	// Ping the database to verify connection
	err = PgDB.Ping(context.Background())
	if err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}

	fmt.Println("Successfully connected to PostgreSQL")

	// dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
	// 				   config.ConnConfig.Host,
	// 				   config.ConnConfig.Port,
	// 				   config.ConnConfig.User,
	// 				   config.ConnConfig.Password,
	//  				   config.ConnConfig.Database,
	// 				   config.)
	GDB, err = gorm.Open(postgres.Open(config.ConnString()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connection established")
	return nil
}

// Close closes the database connection pool.
func Close() {
	if PgDB != nil {
		PgDB.Close()
	}
}
