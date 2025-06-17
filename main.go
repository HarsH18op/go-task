package main

import (
	"log"
	"my-go-task/auth"
	"my-go-task/utils/postgres"
	"my-go-task/watchlist"
	"os"

	"github.com/joho/godotenv"
)

// init() ensures it runs before your main()
func init() {
	// Load env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("❌ Failed to load .env file: %v\n", err)
	} else {
		log.Println("✅ Env varaibles loaded successfully!!")
	}
}

func main() {
	// Connect to database
	db := postgres.ConnectDatabase("configs/postgresDb.yaml")

	// Create a default router
	// router := gin.Default()
	// router.Use(func(c *gin.Context) {
	//     c.Set("db", db) // Setting db in context
	//     c.Next()
	// })

	// port := os.Getenv("SERVER_PORT")
	// router.Run(":" + port)

	go auth.StartAuthServer(db)
	go watchlist.StartWatchlistServer(db)

	authPort := os.Getenv("AUTH_PORT")
	watchlistPort := os.Getenv("WATCHLIST_PORT")
	log.Printf("Auth service started on %v & Watchlist service started on %v port!!", authPort, watchlistPort)

	select {} // To keep main go-routine alive
}
