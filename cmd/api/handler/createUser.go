package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/walber-vaz/crud-go/cmd/api/handler/request"
	"github.com/walber-vaz/crud-go/configs/rest_error"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_error.NewBadRequestError(
			fmt.Sprintf(
				"There are some invalid fields in your request, error: %s", err.Error(),
			),
		)

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
