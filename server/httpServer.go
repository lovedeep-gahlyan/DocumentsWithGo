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
	config             *viper.Viper
	router             *gin.Engine
	dbHandler          *gorm.DB
	userController     *controllers.UserHandler
	documentController *controllers.DocumentHandler
}

func InitHttpServer(config *viper.Viper, dbHandler *gorm.DB) HttpServer {

	userRepository := repositories.NewUserRepository(dbHandler)
	userService := service.NewUserService(userRepository)
	userController := controllers.NewUserHandler(userService)

	documentRepository := repositories.NewDocumentRepository(dbHandler)
	documentService := service.NewDocumentService(documentRepository, userRepository)
	documentController := controllers.NewDocumentHandler(documentService)

	router := gin.Default()

	router.POST("/users", userController.CreateUser)
	router.GET("/users/:user_id", userController.GetUser)
	router.POST("/users/:user_id/documents", documentController.CreateDocument)
	router.PUT("/users/:user_id/documents/:doc_id", documentController.EditDocument)
	router.DELETE("/users/:user_id/documents/:doc_id", documentController.DeleteDocument)
	router.PUT("/users/:user_id/documents/:doc_id/grant/:target_user_id", documentController.GrantAccess)
	router.GET("/users/:user_id/documents/:doc_id", documentController.GetDocument)

	return HttpServer{
		config:             config,
		router:             router,
		dbHandler:          dbHandler,
		userController:     userController,
		documentController: documentController,
	}
}

func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Error while starting HTTP server: %v", err)
	}
}
