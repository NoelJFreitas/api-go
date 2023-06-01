package model

import (
	"github.com/NoelJFreitas/api-go/src/configuration/logger"
	"github.com/NoelJFreitas/api-go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (*UserDomain) CreateUser(UserDomain) *rest_err.RestErr {
	logger.Info("Init createdUser model", zap.String("journey", "createdUser"))

	return nil
}
