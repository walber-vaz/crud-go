package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/walber-vaz/crud-go/cmd/api/handler/request"
	"github.com/walber-vaz/crud-go/cmd/api/model"
	"github.com/walber-vaz/crud-go/configs/logger"
	"github.com/walber-vaz/crud-go/configs/validation"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	logger.Info("Creating user",
		zap.String("path", c.Request.URL.Path),
		zap.String("method", c.Request.Method),
		zap.String("ip", c.ClientIP()),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error binding JSON", err,
			zap.String("path", c.Request.URL.Path),
			zap.String("method", c.Request.Method),
		)
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.FirstName,
		userRequest.LastName,
		userRequest.Age,
	)

	if err := domain.Create(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created",
		zap.String("path", c.Request.URL.Path),
		zap.String("method", c.Request.Method),
		zap.String("ip", c.ClientIP()),
	)

	c.JSON(http.StatusCreated, "User created successfully")
}
