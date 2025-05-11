package main

import (
	"fmt"
	"log"
	"nextgen/internals/config"
	"nextgen/internals/db/ora"
	"nextgen/internals/db/redis"
	"nextgen/internals/db/pg"
	"nextgen/pkg/accounts"
	"nextgen/pkg/auth"
	"nextgen/pkg/server"
	"nextgen/pkg/web/aboutus"
	"nextgen/pkg/web/blog"
	"nextgen/pkg/web/dashboard/team"
)

func init() {
	oraCfg, err := config.LoadOracleConfig()
	if err != nil {
		log.Fatalf("Failed to load oracle database configuration: %v", err)
	}
	oraDsn := fmt.Sprintf(`user="%s" password="%s" connectString="%s" libDir="%s"`,
		oraCfg.OracleUser, oraCfg.OraclePassword, oraCfg.OracleConnectString, oraCfg.OracleInstantClientPath)
	fmt.Println(oraDsn)
	// Initialize the Oracle database connection pool
	if err := ora.Initialize(oraDsn, oraCfg.OracleMaxOpenConns, oraCfg.OracleMaxIdleConns); err != nil {
		log.Fatalf("Oracle database initialization failed: %v", err)
	}

	redisCfg, err := config.LoadRedisConfig()
	if err != nil {
		log.Fatalf("Failed to load redis configuration: %v", err)
	}
      

	// Initialize the Redis connection pool
	redis.Initialize(*redisCfg)


	pgCfg, err := config.LoadPostgresConfig()
	if err != nil {
		log.Fatalf("Failed to lead postgres database configration: %v", err)
	}
	pgDsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		pgCfg.PostgresHost, pgCfg.PostgresPort, pgCfg.PostgresUser, pgCfg.PostgresPassword, pgCfg.PostgresDbName, pgCfg.PostgresSslMode)

	if err := pg.Initialize(pgDsn, pgCfg.PostgresMaxOpenConns, pgCfg.PostgresMaxIdleConns); err != nil {
		log.Fatalf("Postgres database initialization failed: %v", err)
	}

}

func main() {
	// startTime := time.Now()
	// ora.CallProcedure()
	// elapsedTime := time.Since(startTime)
	// fmt.Printf("Execution Time: %s\n", elapsedTime)
	// Create a new DB connection
	err := pg.GDB.AutoMigrate(auth.User{}, blog.BlogPost{}, aboutus.TeamInfo{}, accounts.Memberships{}, team.Employee{}, team.Task{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migrated successfully")

	srv := server.NewServer()

	srv.StartServer()
}
