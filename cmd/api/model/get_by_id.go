package model

import (
	"github.com/walber-vaz/crud-go/configs/logger"
	"github.com/walber-vaz/crud-go/configs/rest_error"
	"go.uber.org/zap"
)

func (u *UserDomain) GetById(id string) (*UserDomain, *rest_error.RestError) {
	logger.Info("Init get by id user model", zap.String("journey", "get_by_id_user"))

	return nil, nil
}
