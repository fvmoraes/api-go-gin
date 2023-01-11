package servers

import (
	"api-go-gin/routes"

	"github.com/gin-gonic/gin"
)

func RunMyFoobarServer() {
	r := gin.Default()
	routes.MyFoobarRoutes(r)
	r.Run(":9000")
}
