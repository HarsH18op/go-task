package auth

import (
	services "my-go-task/auth/business"
	"my-go-task/auth/handlers"
	"my-go-task/auth/repositories"
	routes "my-go-task/auth/router"
	"os"

	_ "my-go-task/auth/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// func init() {
// 	// Load env file
// 	if err := godotenv.Load("../.env"); err != nil {
// 		log.Fatalf("❌ Failed to load .env file: %v\n", err)
// 	} else {
// 		log.Println("✅ Env varaibles loaded successfully!!")
// 	}
// }

// @title User API
// @version 1.0
// @description This is a sample server for managing users.
// @host localhost:8000
// @BasePath /
func StartAuthServer(db *gorm.DB) {
	// func main() {
	// 	// Connect to database
	// 	db := postgres.ConnectDatabase("../configs/postgresDb.yaml")
	router := gin.Default()
	authGroup := router.Group("/api/auth")

	userRepo1 := repositories.NewGetUserRepository(db)
	userRepo2 := repositories.NewCreateUserRepository(db)
	userService1 := services.NewGetUserService(userRepo1)
	userService2 := services.NewCreateUserService(userRepo2)
	userHandler1 := handlers.NewGetUserHandler(userService1)
	userHandler2 := handlers.NewCreateUserHandler(userService2)

	routes.UserApiRoutes(authGroup, userHandler1, userHandler2)
	authGroup.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler, ginSwagger.InstanceName("auth")))
	router.Run(":" + os.Getenv("AUTH_PORT"))
}
