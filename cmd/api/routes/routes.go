package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/walber-vaz/crud-go/cmd/api/handler"
)

func InitRoutes(r *gin.RouterGroup) {
	basePath := "/api/v1"
	v1 := r.Group(basePath)
	{
		v1.GET("/users", handler.GetUsers)
		v1.POST("/user", handler.CreateUser)
		v1.GET("/user/email/:email", handler.FindUserByEmail)
		v1.GET("/user/id/:id", handler.FindUserById)
		v1.PATCH("/user/:id", handler.UpdateUser)
		v1.DELETE("/user/:id", handler.DeleteUser)
	}
}
