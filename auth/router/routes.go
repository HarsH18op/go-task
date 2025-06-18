package routes

import (
	constants "my-go-task/auth/commons"
	"my-go-task/auth/handlers"

	"github.com/gin-gonic/gin"
)

func UserApiRoutes(rg *gin.RouterGroup, handler *handlers.UserHandler) {
	rg.GET(constants.USER_ROUTES.GET_USERS, handler.GetUsers)
	rg.POST(constants.USER_ROUTES.CREATE_USER, handler.CreateUser)
}
