package routes

import (
	"library_app/config"
	"library_app/internal/controller"
	"library_app/manager"
	"log"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) error {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	infra, err := manager.NewInfraManager(cfg)
	if err != nil {
		log.Fatal(err)
	}

	repo := manager.NewRepoManager(infra)
	sv := manager.NewServiceManager(repo)

	authController := controller.NewAuthController(sv.AuthService())

	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", authController.Create)
		authRoutes.POST("/login", authController.Login)
	}

	return router.Run()
}
