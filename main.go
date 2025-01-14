package main

import (
	"log"

	logger "github.com/ForFun21/crud-go/src/configuration/Logger"
	"github.com/ForFun21/crud-go/src/controller"
	"github.com/ForFun21/crud-go/src/controller/routes"
	"github.com/ForFun21/crud-go/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start the application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Init dependencies
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
