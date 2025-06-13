package main

import (
	"log"
	"my-go-task/auth/handlers"
	"my-go-task/auth/repositories"
	"my-go-task/auth/services"
	userRoutes "my-go-task/routes"
	"my-go-task/utils/postgres"
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

	// Create a default router
	router := gin.Default()
	// router.Use(func(c *gin.Context) {
	//     c.Set("db", db) // Setting db in context
	//     c.Next()
	// })

	// Define a route
	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Grouped API routes
	api := router.Group("/api")
	{
		userRoutes.UserApiRoutes(api, userHandler1, userHandler2)
	}

	port := os.Getenv("SERVER_PORT")

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":" + port)
}
