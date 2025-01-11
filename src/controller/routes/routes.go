package routes

import (
	"github.com/ForFun21/crud-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getUserByID/:userID", controller.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/creatUser/", controller.CreatUser)
	r.PUT("/updateUser/:userID", controller.UpdateUser)
	r.DELETE("/deleteUserByID/:userID", controller.DeleteUser)
}
