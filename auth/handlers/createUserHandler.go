package handlers

import (
	"log"
	services "my-go-task/auth/business"
	constants "my-go-task/auth/commons"
	"my-go-task/auth/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// Struct that holds a reference to the service.
type CreateUserHandler struct {
	service *services.CreateUserService
}

// Constructor to create a handler using the given service.
func NewCreateUserHandler(service *services.CreateUserService) *CreateUserHandler {
	return &CreateUserHandler{service: service}
}

// @Summary Create a new user
// @Description Creates a user with provided information
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequestModel true "User data"
// @Success 201 {object} models.CreateUserResponseModel
// @Failure 400 {object} models.ErrorResponseModel
// @Failure 500 {object} models.ErrorResponseModel
// @Router /api/auth/users [post]
func (h *CreateUserHandler) CreateUser(c *gin.Context) {
	var userRequestData models.CreateUserRequestModel
	if err := c.ShouldBind(&userRequestData); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponseModel{
			Message: constants.USER_CREATION_ERRORS.InvalidInputError.Error(),
			// Errors:  map[string]string{"body": err.Error()},
		})
		return
	}

	// This is required for validating data based on validations given in request struct
	if err := validate.Struct(userRequestData); err != nil {
		log.Println(err.Error())
		errorMap := make(map[string]string)
		for _, fieldErr := range err.(validator.ValidationErrors) {
			// errorMap[fieldErr.Field()] = fieldErr.Error() // For showing default error message from validator
			field := fieldErr.Field() // To access value of validation given
			param := fieldErr.Param()
			switch fieldErr.Tag() {
			case "required":
				errorMap[field] = "This field is required"
			case "len":
				errorMap[field] = "Length should be " + param
			case "min":
				errorMap[field] = "Too short"
			case "max":
				errorMap[field] = "Too long"
			case "gte":
				errorMap[field] = "Must be greater than " + param
			case "lte":
				errorMap[field] = "Must be less than or equal to " + param
			case "datetime":
				errorMap[field] = "Must be a valid date (YYYY-MM-DD)"
			default:
				errorMap[field] = "Invalid value"
			}
		}
		c.JSON(http.StatusBadRequest, models.ErrorResponseModel{
			Message: constants.USER_CREATION_ERRORS.ValidationFailedError.Error(),
			Errors:  errorMap,
		})
		return
	}

	// user, err := services.CreateUserService(userRequestData) // If implemented in normal way
	user, err := h.service.CreateUserService(userRequestData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponseModel{
			Message: constants.USER_CREATION_ERRORS.InternalServerError.Error(),
			Errors:  map[string]string{"server": err.Error()},
		})
		return
	}

	// Convert DB models to response format
	response := models.CreateUserResponseModel{
		ID:      user.ID,
		Message: constants.USER_CREATION_SUCCESS_MESSAGE,
	}
	c.JSON(http.StatusCreated, response)
}

// If implemented in normal way
// func GetUsers(c *gin.Context) {
// 	users, err := services.GetAllUsersService()
