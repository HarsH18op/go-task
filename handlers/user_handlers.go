package handlers

import (
	"my-go-task/models"
	"my-go-task/services"
	my_validators "my-go-task/utils/validators"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Struct that holds a reference to the service.
type UserHandler struct {
	service *services.UserService
}

// Constructor to create a handler using the given service.
func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// This is the actual HTTP GET handler function.
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.service.GetAllUsersService()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func (h *UserHandler) CreateUser(c *gin.Context){
	var userRequestData models.CreateUserRequest
	if err := c.ShouldBind(&userRequestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	// This is required for validating data based on validations given in request struct
	if err := my_validators.Validate.Struct(userRequestData); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errors := make(map[string]string)
		for _, fieldErr := range validationErrors {
			errors[fieldErr.Field()] = my_validators.ValidationMessage(fieldErr)
		}
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": errors})
		return
	}

	// user, err := services.CreateUserService(userRequestData) // If implemented in normal way
	user, err := h.service.CreateUserService(userRequestData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create user",
			"details": err.Error(), // actual DB error message
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully.", "details": user})
}

// If implemented in normal way
// func GetUsers(c *gin.Context) {
// 	users, err := services.GetAllUsersService()