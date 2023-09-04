package model

import (
	"github.com/walber-vaz/crud-go/configs/logger"
	"github.com/walber-vaz/crud-go/configs/rest_error"
	"go.uber.org/zap"
)

func (u *UserDomain) GetAll() ([]UserDomain, *rest_error.RestError) {
	logger.Info("Init get all user model", zap.String("journey", "get_all_user"))

	return nil, nil
}
