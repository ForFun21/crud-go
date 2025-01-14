package controller

import (
	"net/http"

	logger "github.com/ForFun21/crud-go/src/configuration/Logger"
	"github.com/ForFun21/crud-go/src/configuration/rest_err/validation"
	"github.com/ForFun21/crud-go/src/controller/model/request"
	"github.com/ForFun21/crud-go/src/model"
	"github.com/ForFun21/crud-go/src/model/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
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

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)
	service := service.NewUserDomainService()
	if err := service.CreatUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "CreateUser"))

	c.String(http.StatusOK, "")
}
