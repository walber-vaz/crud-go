package model

import (
	"fmt"

	"github.com/walber-vaz/crud-go/configs/logger"
	"github.com/walber-vaz/crud-go/configs/rest_error"
	"go.uber.org/zap"
)

func (u *UserDomain) Create() *rest_error.RestError {
	logger.Info("Init create user model", zap.String("journey", "create_user"))

	u.EncryptPassword()

	fmt.Println(u.Password)

	return nil
}
