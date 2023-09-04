package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/walber-vaz/crud-go/cmd/api/handler/request"
	"github.com/walber-vaz/crud-go/cmd/api/handler/response"
	"github.com/walber-vaz/crud-go/configs/validation"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error when trying to bind user request: %s", err.Error())
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
	response := response.UserResponse{
		ID:        "1",
		Email:     userRequest.Email,
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
		Age:       userRequest.Age,
	}

	c.JSON(http.StatusCreated, response)
}
