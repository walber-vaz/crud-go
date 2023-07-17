package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/walber-vaz/backend-go-gin/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	base_path := "/api/v1"
	v1 := r.Group(base_path)
	{
		v1.GET("/users/:id", controller.FindUser)
		v1.PUT("/users/:id", controller.UpdateUser)
		v1.DELETE("/users/:id", controller.DeleteUser)
		v1.POST("/users", controller.CreateUser)
		v1.GET("/users", controller.FindAllUsers)
	}
}
