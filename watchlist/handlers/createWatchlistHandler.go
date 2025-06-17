package handlers

import (
	"log"
	constants "my-go-task/watchlist/commons"
	"my-go-task/watchlist/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// @Summary Create a new watchlist
// @Description Creates a watchlist with provided information
// @Tags Watchlists
// @Accept json
// @Produce json
// @Param watchlist body models.CreateWatchlistRequestModel true "Watchlist data"
// @Success 201 {object} models.CreateWatchlistResponseModel
// @Failure 400 {object} models.ErrorResponseModel
// @Failure 500 {object} models.ErrorResponseModel
// @Router /api/watchlists [post]
func (h *WatchlistHandler) CreateWatchlist(c *gin.Context) {
	var watchlistRequestData models.CreateWatchlistRequestModel
	if err := c.ShouldBind(&watchlistRequestData); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponseModel{
			Message: constants.WATCHLIST_CREATION_ERRORS.InvalidInputError.Error(),
			// Errors:  map[string]string{"body": err.Error()},
		})
		return
	}

	// This is required for validating data based on validations given in request struct
	if err := validate.Struct(watchlistRequestData); err != nil {
		errorMap := make(map[string]string)
		for _, fieldErr := range err.(validator.ValidationErrors) {
			field := fieldErr.Field() // To access value of validation given
			switch fieldErr.Tag() {
			case "required":
				errorMap[field] = "This field is required"
			default:
				errorMap[field] = "Invalid value"
			}
		}
		c.JSON(http.StatusBadRequest, models.ErrorResponseModel{
			Message: constants.WATCHLIST_CREATION_ERRORS.ValidationFailedError.Error(),
			Errors:  errorMap,
		})
		return
	}

	watchlist, err := h.service.CreateWatchlistService(watchlistRequestData)
	if err != nil {
		log.Println(err.Error())
		if constants.IsClientError(err) {
			c.JSON(http.StatusBadRequest, models.ErrorResponseModel{Message: err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, models.ErrorResponseModel{Message: constants.WATCHLIST_CREATION_ERRORS.InternalServerError.Error()})
		}
		return
	}
	// Convert DB models to response format
	response := models.CreateWatchlistResponseModel{
		ID:      watchlist.ID,
		Message: constants.WATCHLIST_CREATION_SUCCESS_MESSAGE,
	}
	c.JSON(http.StatusCreated, response)
}
