package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/walber-vaz/backend-go-gin/util/rest_errors"
)

func CreateUser(ctx *gin.Context) {
	err := rest_errors.NewBadRequestError("invalid json body")
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(err.StatusCode, err)
}
