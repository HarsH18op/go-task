package watchlist

import (
	services "my-go-task/watchlist/business"
	"my-go-task/watchlist/handlers"
	"my-go-task/watchlist/repositories"
	routes "my-go-task/watchlist/router"
	"os"

	_ "my-go-task/watchlist/docs"

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

// @title Watchlist API
// @version 1.0
// @description This is a sample server for managing watchlists.
// @host localhost:8082
// @BasePath /
func StartWatchlistServer(db *gorm.DB) {
	// func main() {
	// 	// Connect to database
	// 	db := postgres.ConnectDatabase("../configs/postgresDb.yaml")
	router := gin.Default()
	watchlistGroup := router.Group("/api/watchlist")
	watchlistRepo := repositories.NewWatchlistRepository(db)
	watchlistService := services.NewWatchlistService(watchlistRepo)
	watchlistHandler := handlers.NewWatchlistHandler(watchlistService)

	routes.WatchlistApiRoutes(watchlistGroup, watchlistHandler)
	watchlistGroup.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler, ginSwagger.InstanceName("watchlist")))
	router.Run(":" + os.Getenv("WATCHLIST_PORT"))
}
