package routes

import (
	"CodeSheep-runcode/controllers"
	"CodeSheep-runcode/middles"

	"github.com/gin-gonic/gin"
)

var (
	userController controllers.CodeController
)

func init() {
	userController = controllers.CodeController{}
}

func Start(port string) {
	router := gin.Default()
	router.Use(middles.SetSessionId())

	router.POST("/code-run", userController.HandleRunCode)

	router.Run(port)
}
