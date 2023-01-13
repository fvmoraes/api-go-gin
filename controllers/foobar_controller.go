package controllers

import (
	"api-go-gin/helpers"
	"api-go-gin/initializers"
	"api-go-gin/logs"
	"api-go-gin/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateFoobar(c *gin.Context) {
	var foobar models.Foobar
	if err := c.ShouldBindJSON(&foobar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		logs.PopulateErrorLogFile("ERROR", err, "Bind To Json Error")
		return
	}
	if err := helpers.ModelValidator(&foobar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		logs.PopulateErrorLogFile("WARNING", err, "Data validation with errors")
		return
	}
	initializers.DB.Create(&foobar)
	c.JSON(http.StatusOK, foobar)
	logs.PopulateErrorLogFile("INFO", nil, "Successful call to endpoint: "+""+c.Request.Method+""+c.Request.RequestURI)
}

func ShownFoobar(c *gin.Context) {
	var foobar []models.Foobar
	if err := initializers.DB.Find(&foobar); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": err.Error.Error(),
		})
		logs.PopulateErrorLogFile("ERROR", err.Error, "Find Error")
		return
	}
	c.JSON(http.StatusOK, foobar)
	logs.PopulateErrorLogFile("INFO", nil, "Successful call to endpoint: "+""+c.Request.Method+""+c.Request.RequestURI)
}

func ShownFoobarByParamId(c *gin.Context) {
	var foobar models.Foobar
	id := c.Params.ByName("id")
	if err := initializers.DB.First(&foobar, id); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": err.Error.Error(),
		})
		logs.PopulateErrorLogFile("ERROR", err.Error, "Find Error")
		return
	}
	c.JSON(http.StatusOK, foobar)
	logs.PopulateErrorLogFile("INFO", nil, "Successful call to endpoint: "+""+c.Request.Method+""+c.Request.RequestURI)
}

func EditFoobarByParamId(c *gin.Context) {
	var foobar models.Foobar
	id := c.Params.ByName("id")
	if err := initializers.DB.First(&foobar, id); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": err.Error.Error(),
		})
		logs.PopulateErrorLogFile("ERROR", err.Error, "Find Error")
		return
	}
	if err := c.ShouldBindJSON(&foobar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		logs.PopulateErrorLogFile("ERROR", err, "Bind To Json Error")
		return
	}
	if err := helpers.ModelValidator(&foobar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		logs.PopulateErrorLogFile("WARNING", err, "Data validation with errors")
		return
	}
	initializers.DB.Model(&foobar).Updates(foobar)
	c.JSON(http.StatusOK, foobar)
	logs.PopulateErrorLogFile("INFO", nil, "Successful call to endpoint: "+""+c.Request.Method+""+c.Request.RequestURI)
}

func DeleteFoobarByParamId(c *gin.Context) {
	var foobar models.Foobar
	id := c.Params.ByName("id")
	if err := initializers.DB.First(&foobar, id); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": err.Error.Error(),
		})
		logs.PopulateErrorLogFile("ERROR", err.Error, "Find Error")
		return
	}
	initializers.DB.Delete(&foobar, id)
	c.JSON(http.StatusAccepted, foobar)
	logs.PopulateErrorLogFile("INFO", nil, "Successful call to endpoint: "+""+c.Request.Method+""+c.Request.RequestURI)
}

func ShownFoobarByParamReg(c *gin.Context) {
	var foobar []models.Foobar
	reg, _ := strconv.ParseUint(c.Params.ByName("reg"), 10, 64)
	if err := initializers.DB.Where(&models.Foobar{Registration: reg}).Find(&foobar); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": err.Error.Error(),
		})
		logs.PopulateErrorLogFile("ERROR", err.Error, "Find Error")
		return
	}
	c.JSON(http.StatusOK, foobar)
	logs.PopulateErrorLogFile("INFO", nil, "Successful call to endpoint: "+""+c.Request.Method+""+c.Request.RequestURI)
}

func ShownFoobarMockToLearnTests(c *gin.Context) {
	//Mock reference
	c.JSON(200, helpers.MyFoobar)
	logs.PopulateErrorLogFile("INFO", nil, "Successful call to endpoint: "+""+c.Request.Method+""+c.Request.RequestURI)
}
