package model

import (
	"github.com/walber-vaz/crud-go/configs/logger"
	"github.com/walber-vaz/crud-go/configs/rest_error"
	"go.uber.org/zap"
)

func (u *UserDomain) GetByEmail(email string) (*UserDomain, *rest_error.RestError) {
	logger.Info("Init get by email user model", zap.String("journey", "get_by_email_user"))

	return nil, nil
}
