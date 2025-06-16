package handlers

import (
	"log"
	"my-go-task/watchlist/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// @Summary Create a new user
// @Description Creates a user with provided information
// @Tags Watchlists
// @Accept json
// @Produce json
// @Param user body models.CreateWatchlistRequestModel true "Watchlist data"
// @Success 201 {object} models.CreateWatchlistResponseModel
// @Failure 400 {object} models.ErrorResponseModel
// @Failure 500 {object} models.ErrorResponseModel
// @Router /api/users [post]
func (h *WatchlistHandler) CreateWatchlist(c *gin.Context) {
	var watchlistRequestData models.CreateWatchlistRequestModel
	if err := c.ShouldBind(&watchlistRequestData); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponseModel{
			Message: "Invalid input",
			Errors:  map[string]string{"body": err.Error()},
		})
		return
	}

	// This is required for validating data based on validations given in request struct
	if err := validate.Struct(watchlistRequestData); err != nil {
		errorMap := make(map[string]string)
		for _, fieldErr := range err.(validator.ValidationErrors) {
			field := fieldErr.Field() // To access value of validation given
			// param := fieldErr.Param()
			switch fieldErr.Tag() {
			case "required":
				errorMap[field] = "This field is required"
			default:
				errorMap[field] = "Invalid value"
			}
		}
		c.JSON(http.StatusBadRequest, models.ErrorResponseModel{
			Message: "Validation failed",
			Errors:  errorMap,
		})
		return
	}

	watchlist, err := h.service.CreateWatchlistService(watchlistRequestData)
	log.Println(watchlist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponseModel{
			Message: "Failed to create watchlist",
			Errors:  map[string]string{"server": err.Error()},
		})
		return
	}

	// Convert DB models to response format
	response := models.CreateWatchlistResponseModel{
		ID:          watchlist.ID,
		Name:        watchlist.Name,
		UserId:      watchlist.UserID,
		ScriptCount: watchlist.ScriptCount,
		CreatedAt:   watchlist.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   watchlist.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Watchlist created successfully.", "details": response})
}
