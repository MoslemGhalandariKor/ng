package main

import (
	"log"
	"nextgen/internals/db/pg"
	"nextgen/pkg/accounts"
	"nextgen/pkg/auth"
	"nextgen/pkg/server"
	"nextgen/pkg/web/aboutus"
	"nextgen/pkg/web/blog"
	"nextgen/pkg/web/dashboard/team"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	pg.ConnectDatabase()
	pg.Initialize()

}

func main() {

	// Create a new DB connection
	err := pg.GDB.AutoMigrate(auth.User{}, blog.BlogPost{}, aboutus.TeamInfo{}, accounts.Memberships{}, team.Employee{}, team.Task{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migrated successfully")

	srv := server.NewServer()

	srv.StartServer()
}
