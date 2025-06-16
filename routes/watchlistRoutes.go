package routes

import (
	"my-go-task/watchlist/handlers"

	"github.com/gin-gonic/gin"
)

func WatchlistApiRoutes(rg *gin.RouterGroup, handler *handlers.WatchlistHandler) {
	// rg.GET("/users", handler.GetUsers)
	rg.POST("/watchlists", handler.CreateWatchlist)
}
