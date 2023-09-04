package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/walber-vaz/crud-go/cmd/api/handler/request"
	"github.com/walber-vaz/crud-go/cmd/api/handler/response"
	"github.com/walber-vaz/crud-go/configs/rest_error"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error when trying to bind user request: %s", err.Error())
		restErr := rest_error.NewBadRequestError("Some fields are not valid")

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
