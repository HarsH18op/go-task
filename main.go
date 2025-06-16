package main

import (
	"log"
	"my-go-task/auth/handlers"
	"my-go-task/auth/repositories"
	userRouter "my-go-task/auth/router"
	"my-go-task/auth/services"
	"my-go-task/utils/postgres"
	watchlistHandler "my-go-task/watchlist/handlers"
	watchlistRepo "my-go-task/watchlist/repositories"
	watchlistRouter "my-go-task/watchlist/router"
	watchlistService "my-go-task/watchlist/services"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "my-go-task/docs"

	// Replace with your module name

	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerFiles "github.com/swaggo/files"
)

// init() ensures it runs before your main()
func init() {
	// Load env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("❌ Failed to load .env file: %v\n", err)
	} else {
		log.Println("Env varaibles loaded successfully!!")
	}
}

// @title User API
// @version 1.0
// @description This is a sample server for managing users.
// @host localhost:8000
// @BasePath /
func main() {
	// Connect to database
	db := postgres.ConnectDatabase()

	userRepo1 := repositories.NewGetUserRepository(db)
	userRepo2 := repositories.NewCreateUserRepository(db)
	userService1 := services.NewGetUserService(userRepo1)
	userService2 := services.NewCreateUserService(userRepo2)
	userHandler1 := handlers.NewGetUserHandler(userService1)
	userHandler2 := handlers.NewCreateUserHandler(userService2)

	watchlistRepo := watchlistRepo.NewWatchlistRepository(db)
	watchlistService := watchlistService.NewWatchlistService(watchlistRepo)
	watchlistHandler := watchlistHandler.NewWatchlistHandler(watchlistService)

	// Create a default router
	router := gin.Default()
	// router.Use(func(c *gin.Context) {
	//     c.Set("db", db) // Setting db in context
	//     c.Next()
	// })

	// Grouped API routes
	api := router.Group("/api")
	{
		userRouter.UserApiRoutes(api, userHandler1, userHandler2)
		watchlistRouter.WatchlistApiRoutes(api, watchlistHandler)
	}

	port := os.Getenv("SERVER_PORT")

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":" + port)
}
