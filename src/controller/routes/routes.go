package routes

import (
	"github.com/ForFun21/crud-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface) {

	r.GET("/getUserByID/:userID", userController.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	r.POST("/creatUser/", userController.CreatUser)
	r.PUT("/updateUser/:userID", userController.UpdateUser)
	r.DELETE("/deleteUserByID/:userID", userController.DeleteUser)
}
