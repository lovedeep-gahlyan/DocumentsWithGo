package main

import (
	"docs/config"
	"docs/server"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Println("Starting Mutual Fund App")

	// taking docs.toml file
	log.Println("Initializig configuration")
	config := config.InitConfig("docs")

	log.Println("Initializig database")
	dbHandler := server.InitDatabase(config)

	log.Println("Initializig HTTP sever")
	httpServer := server.InitHttpServer(config, dbHandler)

	httpServer.Start()
}
