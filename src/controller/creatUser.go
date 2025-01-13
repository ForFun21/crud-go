package controller

import (
	"net/http"

	logger "github.com/ForFun21/crud-go/src/configuration/Logger"
	"github.com/ForFun21/crud-go/src/configuration/rest_err/validation"
	"github.com/ForFun21/crud-go/src/controller/model/request"
	"github.com/ForFun21/crud-go/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "CreateUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "CreateUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	logger.Info("User created successfully",
		zap.String("journey", "CreateUser"))

	c.JSON(http.StatusOK, response)
}
