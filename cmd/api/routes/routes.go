package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/walber-vaz/crud-go/cmd/api/handler"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/users", handler.GetUsers)
	r.POST("/user", handler.CreateUser)
	r.GET("/user/email/:email", handler.FindUserByEmail)
	r.GET("/user/id/:id", handler.FindUserById)
	r.PATCH("/user/:id", handler.UpdateUser)
	r.DELETE("/user/:id", handler.DeleteUser)
}
