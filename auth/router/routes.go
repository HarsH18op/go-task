package routes

import (
	"my-go-task/auth/handlers"

	"github.com/gin-gonic/gin"
)

func UserApiRoutes(rg *gin.RouterGroup, handler *handlers.GetUserHandler, handler2 *handlers.CreateUserHandler) {
	rg.GET("/users", handler.GetUsers)
	rg.POST("/users", handler2.CreateUser)
}
