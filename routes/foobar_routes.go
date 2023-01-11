package routes

import (
	"api-go-gin/controllers"

	"github.com/gin-gonic/gin"
)

func MyFoobarRoutes(r *gin.Engine) {
	g := r.Group("/api/v1")
	g.GET("/foobar", controllers.ShownFoobar)
	g.GET("/foobar/:id", controllers.ShownFoobarByParamId)
	g.POST("/foobar", controllers.CreateFoobar)
	g.PATCH("/foobar/:id", controllers.EditFoobarByParamId)
	g.DELETE("/foobar/:id", controllers.DeleteFoobarByParamId)
	g.GET("/foobar/param/:reg", controllers.ShownFoobarByParamReg)
}
