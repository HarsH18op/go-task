package routes

import (
	"my-go-task/handlers"

	"github.com/gin-gonic/gin"
)

func UserApiRoutes(rg *gin.RouterGroup, handler *handlers.UserHandler) {
	rg.GET("/users", handler.GetUsers)
	rg.POST("/users", handler.CreateUser)
}
