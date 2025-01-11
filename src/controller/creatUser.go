package controller

import (
	"github.com/ForFun21/crud-go/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func CreatUser(c *gin.Context) {

	err := rest_err.NewBadRequestError("Voce chamou a rota de forma errada.")
	c.JSON(err.Code, err)
}
