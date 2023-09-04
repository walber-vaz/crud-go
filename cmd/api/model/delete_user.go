package model

import (
	"github.com/walber-vaz/crud-go/configs/logger"
	"github.com/walber-vaz/crud-go/configs/rest_error"
	"go.uber.org/zap"
)

func (u *UserDomain) Delete(id string) *rest_error.RestError {
	logger.Info("Init delete user model", zap.String("journey", "delete_user"))

	return nil
}
