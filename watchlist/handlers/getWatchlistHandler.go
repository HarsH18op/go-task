package handlers

import (
	"log"
	"my-go-task/watchlist/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get all watchlists
// @Description Returns a list of all watchlists
// @Tags Watchlists
// @Accept json
// @Produce json
// @Success 200 {object} models.GetWatchlistResponseModel
// @Failure 500 {object} models.ErrorResponseModel
// @Router /api/watchlists [get]
// This is the actual HTTP GET handler function.
func (h *WatchlistHandler) GetWatchlists(c *gin.Context) {
	userIDStr := c.Query("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64) // Converts string query param to an unsigned integer (uint64).
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
		return
	}
	watchlists, err := h.service.GetWatchlistsService(uint(userID)) // This converts the uint64 (returned by ParseUint) to uint
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	log.Println(watchlists, "@@@@@@@@@")

	// Convert DB models to response format
	var response = make([]models.GetWatchlistResponseModel, 0)

	for _, u := range watchlists {
		resp := models.GetWatchlistResponseModel{
			ID:   u.ID,
			Name: u.Name,
		}
		response = append(response, resp)
	}
	c.JSON(http.StatusOK, gin.H{"watchlists": response})
}
