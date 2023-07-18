package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/walber-vaz/backend-go-gin/model/request"
	"github.com/walber-vaz/backend-go-gin/util/rest_errors"
)

func CreateUser(ctx *gin.Context) {
	var userRequest request.UserRequest
	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_errors.NewBadRequestError(
			fmt.Sprintf("The request body is not valid: %s", err.Error()))
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(restErr.StatusCode, restErr)
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, userRequest)
}
