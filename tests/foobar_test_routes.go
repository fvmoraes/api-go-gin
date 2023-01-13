package tests

import "github.com/gin-gonic/gin"

func RouteSetupTesting() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}
