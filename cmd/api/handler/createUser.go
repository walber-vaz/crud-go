package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/walber-vaz/crud-go/configs/rest_error"
)

func CreateUser(c *gin.Context) {
	err := rest_error.NewBadRequestError("Bad request")
	c.JSON(err.Code, err)
}
