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

	v1 := router.Group("/api/v1")

	library := v1.Group("/library")

	authRoutes:= library.Group("/auth")
	{
		authRoutes.POST("/register", authController.CreateUser)
		authRoutes.POST("/login", authController.LoginUser)
	}

	return router.Run()
}
