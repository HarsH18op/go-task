package main

import (
	"log"
	"my-go-task/handlers"
	"my-go-task/repositories"
	"my-go-task/routes"
	"my-go-task/services"
	"my-go-task/utils/postgres"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
	_ "my-go-task/docs"
)

// gin-swagger middleware

// init() ensures it runs before your main()
func init() {
	// Load env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("❌ Failed to load .env file: %v\n", err)
	} else {
		log.Println("Env varaibles loaded successfully!!")
	}
}

func main() {
	// Connect to database
	db := postgres.ConnectDatabase()

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

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
		routes.UserApiRoutes(api, userHandler)
	}

	port := os.Getenv("SERVER_PORT")

	// router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":" + port)
}
