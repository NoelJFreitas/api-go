package controller

import (
	"net/http"

	"github.com/NoelJFreitas/api-go/src/configuration/logger"
	"github.com/NoelJFreitas/api-go/src/configuration/validation"
	"github.com/NoelJFreitas/api-go/src/controller/model/request"
	"github.com/NoelJFreitas/api-go/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("journey", "CreateUser"))

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error tryng to validate user info", err,
			zap.String("journey", "CreateUser"),
		)

		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	logger.Info("User created successfully", zap.String("journey", "CreateUser"))

	response := response.UserResponse{
		ID:    "teste",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	c.JSON(http.StatusOK, response)
}
