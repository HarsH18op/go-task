package handlers

import (
	"my-go-task/auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get all users
// @Description Returns a list of all users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} models.GetUserResponseModel
// @Failure 500 {object} models.ErrorResponseModel
// @Router /api/auth/users [get]
// This is the actual HTTP GET handler function.
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.service.GetAllUsersService()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Convert DB models to response format
	var response []models.GetUserResponseModel

	for _, u := range users {
		var birthday *string
		if u.Birthday != nil {
			b := u.Birthday.Format("2006-01-02")
			birthday = &b
		}

		resp := models.GetUserResponseModel{
			ID:        u.ID,
			Name:      u.Name,
			Email:     u.Email,
			Age:       u.Age,
			Pancard:   u.Pancard,
			Mobile:    u.Mobile,
			Bio:       u.Bio,
			Birthday:  birthday,
			CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
		response = append(response, resp)
	}
	c.JSON(http.StatusOK, gin.H{"data": response})
}
