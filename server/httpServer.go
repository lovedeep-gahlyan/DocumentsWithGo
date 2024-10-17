package server

import (
	"docs/controllers"
	"docs/repositories"
	"docs/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type HttpServer struct {
	config         *viper.Viper
	router         *gin.Engine
	dbHandler      *gorm.DB
	userController *controllers.UserHandler
}

func InitHttpServer(config *viper.Viper, dbHandler *gorm.DB) HttpServer {

	userRepository := repositories.NewUserRepository(dbHandler)
	userService := service.NewUserService(userRepository)
	userController := controllers.NewUserHandler(userService)

	router := gin.Default()

	router.POST("/users", userController.CreateUser)

	return HttpServer{
		config:         config,
		router:         router,
		dbHandler:      dbHandler,
		userController: userController,
	}
}

func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Error while starting HTTP server: %v", err)
	}
}
