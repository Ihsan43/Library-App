package routes

import (
	"library_app/config"
	"library_app/internal/controller"
	"library_app/internal/middleware"
	"library_app/manager"
	"log"

	"github.com/gin-contrib/cors"
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

	// Set CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://192.168.18.4:8081"}, // Frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Setup Routes
	v1 := router.Group("/api/v1")
	library := v1.Group("/library")

	// Auth Routes
	authRoutes := library.Group("/auth")
	{
		authRoutes.POST("/register", authController.CreateUser)
		authRoutes.POST("/login", authController.LoginUser)
	}

	// User Routes with Auth Middleware
	user := library.Group("", middleware.AuthMiddleware(), middleware.ValidationMiddleware())
	{
		user.GET("/user/:id", userController.GetUserId)
		user.PUT("/user/:id", userController.UpdatedUserById)
		user.GET("/users", userController.GetUsersWithPagination)
		user.DELETE("/user/:id", userController.DeleteUserID)
	}

	// Book Routes with Auth Middleware
	bookController := controller.NewBookController(sv.BookService())
	book := library.Group("", middleware.AuthMiddleware())
	{
		book.POST("/book", bookController.CreateBook)
		book.GET("/book/:id", bookController.GetBook)
		book.GET("/books", bookController.GetBooksWithPagination)
		book.PUT("/book/:id", bookController.UpdatedBookById)
		book.DELETE("/book/:id", bookController.DeleteBook)
	}

	// Address Routes with Auth Middleware
	addressController := controller.NewAddressController(sv.AddressService())
	address := library.Group("", middleware.AuthMiddleware())
	{
		address.POST("/address", addressController.CreateAddress)
		address.PUT("/address/:id", addressController.UpdateAddress)
		address.GET("/address/:id", addressController.GetAddress)
		address.DELETE("/address/:id", addressController.DeleteAddress)
	}

	// Order Routes with Auth Middleware
	orderController := controller.NewOrderController(sv.OrderService())
	order := library.Group("", middleware.AuthMiddleware())
	{
		order.POST("/order", orderController.CreateOrder)
	}

	// Payment Routes with Auth Middleware
	paymentController := controller.NewPaymentController(sv.PaymentService())
	payment := library.Group("", middleware.AuthMiddleware())
	{
		payment.POST("/payment", paymentController.CreatePayment)
	}

	// Transaction Routes with Auth Middleware
	transactionController := controller.NewTransactionController(sv.TransactionService())
	transaction := library.Group("", middleware.AuthMiddleware())
	{
		transaction.GET("/transaction", transactionController.GetTransactionHistories)
	}

	// Run server
	return router.Run()
}
