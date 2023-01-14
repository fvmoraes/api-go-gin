package servers

import (
	"api-go-gin/docs"
	"api-go-gin/routes"

	"github.com/gin-gonic/gin"
)

func RunMyFoobarServer() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	routes.MyFoobarRoutes(r)
	r.Run(":9000")
}
