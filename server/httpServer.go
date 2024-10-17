package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type HttpServer struct {
	config    *viper.Viper
	router    *gin.Engine
	dbHandler *gorm.DB
}

func InitHttpServer(config *viper.Viper, dbHandler *gorm.DB) HttpServer {

	router := gin.Default()

	return HttpServer{
		config:    config,
		router:    router,
		dbHandler: dbHandler,
	}
}

func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Error while starting HTTP server: %v", err)
	}
}
