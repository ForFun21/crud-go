package controller

import (
	"github.com/ForFun21/crud-go/src/model/service"
	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(
	serviceInterface service.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	FindUserByID(c *gin.Context)
	FindUserByEmail(c *gin.Context)

	DeleteUser(c *gin.Context)
	CreatUser(c *gin.Context)
	UpdateUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
