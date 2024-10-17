package server

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config *viper.Viper
	router *gin.Engine
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {

	router := gin.Default()

	return HttpServer{
		config: config,
		router: router,
	}
}

func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Error while starting HTTP server: %v", err)
	}
}
