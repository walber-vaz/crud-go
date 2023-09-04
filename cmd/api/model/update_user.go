package model

import (
	"github.com/walber-vaz/crud-go/configs/logger"
	"github.com/walber-vaz/crud-go/configs/rest_error"
	"go.uber.org/zap"
)

func (u *UserDomain) Update(id string) *rest_error.RestError {
	logger.Info("Init update user model", zap.String("journey", "update_user"))

	return nil
}
