package main

import (
	"CodeSheep-runcode/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	routes.Start(":8080")
}
