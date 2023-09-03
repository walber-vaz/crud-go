package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/walber-vaz/crud-go/cmd/api/handler"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/users", handler.GetUsers)
	r.POST("/user", handler.CreateUser)
	r.GET("/user/:id", handler.FindUserById)
	r.GET("/user/:email", handler.FindUserByEmail)
	r.PATCH("/user/:id", handler.UpdateUser)
	r.DELETE("/user/:id", handler.DeleteUser)
}
