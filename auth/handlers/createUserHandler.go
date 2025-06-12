package handlers

import (
	"log"
	"my-go-task/auth/models"
	"my-go-task/auth/services"
	my_validators "my-go-task/utils/validators"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Struct that holds a reference to the service.
type CreateUserHandler struct {
	service *services.CreateUserService
}

// Constructor to create a handler using the given service.
func NewCreateUserHandler(service *services.CreateUserService) *CreateUserHandler {
	return &CreateUserHandler{service: service}
}

func (h *CreateUserHandler) CreateUser(c *gin.Context) {
	var userRequestData models.CreateUserRequestModel
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
	log.Println(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create user",
			"details": err.Error(), // actual DB error message
		})
		return
	}

	// Convert DB models to response format
	var response models.CreateUserResponseModel
	var birthday *string
	if user.Birthday != nil {
		b := user.Birthday.Format("2006-01-02")
		birthday = &b
	}

	response = models.CreateUserResponseModel{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Age:       user.Age,
		Pancard:   user.Pancard,
		Mobile:    user.Mobile,
		Bio:       user.Bio,
		Birthday:  birthday,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully.", "details": response})
}

// If implemented in normal way
// func GetUsers(c *gin.Context) {
// 	users, err := services.GetAllUsersService()
