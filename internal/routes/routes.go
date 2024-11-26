package routes

import (
	"library_app/config"
	"library_app/internal/controller"
	"library_app/internal/middleware"
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

	userController := controller.NewUSerController(sv.UserService())

	v1 := router.Group("/api/v1")

	library := v1.Group("/library")

	authRoutes := library.Group("/auth")
	{
		authRoutes.POST("/register", authController.CreateUser)
		authRoutes.POST("/login", authController.LoginUser)
	}

	user := library.Group("", middleware.AuthMiddleware(), middleware.ValidationMiddleware())

	{
		user.GET("/user/:id", userController.GetUserId)
		user.PUT("/user/:id", userController.UpdatedUserById)
		user.GET("/users", userController.GetUsersWithPagination)
		user.DELETE("/user/:id", userController.DeleteUserID)
	}

	bookController := controller.NewBookController(sv.BookService())

	book := library.Group("", middleware.AuthMiddleware())

	{
		book.POST("/book", bookController.CreateBook)
		book.GET("/book/:id", bookController.GetBook)
		book.GET("/books", bookController.GetBooksWithPagination)
		book.DELETE("/book/:id", bookController.DeleteBook)
	}

	addressController := controller.NewAddressController(sv.AddressService())

	address := library.Group("", middleware.AuthMiddleware(), middleware.AuthMiddleware())

	{
		address.POST("/address", addressController.CreateAddress)
		address.PUT("/address/:id", addressController.UpdateAddress)
		address.GET("/address/:id", addressController.GetAddress)
		address.DELETE("/address/:id", addressController.DeleteAddress)
	}

	return router.Run()
}
