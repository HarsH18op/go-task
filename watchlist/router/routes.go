package routes

import (
	constants "my-go-task/watchlist/commons"
	"my-go-task/watchlist/handlers"

	"github.com/gin-gonic/gin"
)

func WatchlistApiRoutes(rg *gin.RouterGroup, handler *handlers.WatchlistHandler) {
	rg.GET(constants.WATCHLIST_ROUTES.GET_WATCHLISTS, handler.GetWatchlists)
	rg.POST(constants.WATCHLIST_ROUTES.CREATE_WATCHLIST, handler.CreateWatchlist)
}
