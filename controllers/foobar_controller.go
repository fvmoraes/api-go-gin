package controllers

import (
	"api-go-gin/helpers"
	"api-go-gin/initializers"
	"api-go-gin/models"
	"log"
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
		log.Panic("Bind To Json Error", err)
		return
	}
	if err := helpers.ModelValidator(&foobar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		log.Panic("Data validation with errors", err)
		return
	}
	initializers.DB.Create(&foobar)
	c.JSON(http.StatusOK, foobar)
}

func ShownFoobar(c *gin.Context) {
	var foobar []models.Foobar
	//Mock reference removed
	//c.JSON(200, helpers.MyFoobar)
	if err := initializers.DB.Find(&foobar); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": err.Error.Error(),
		})
		log.Panic("Find Error", err.Error.Error())
		return
	}
	c.JSON(http.StatusOK, foobar)
}

func ShownFoobarByParamId(c *gin.Context) {
	var foobar models.Foobar
	id := c.Params.ByName("id")
	if err := initializers.DB.First(&foobar, id); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": err.Error.Error(),
		})
		log.Panic("Find Error", err.Error.Error())
		return
	}
	c.JSON(http.StatusOK, foobar)
}

func EditFoobarByParamId(c *gin.Context) {
	var foobar models.Foobar
	id := c.Params.ByName("id")
	if err := initializers.DB.First(&foobar, id); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": err.Error.Error(),
		})
		log.Panic("Find Error", err.Error.Error())
		return
	}
	if err := c.ShouldBindJSON(&foobar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		log.Panic("Bind To Json Error", err)
		return
	}
	if err := helpers.ModelValidator(&foobar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		log.Panic("Data validation with errors", err)
		return
	}
	initializers.DB.Model(&foobar).Updates(foobar)
	c.JSON(http.StatusOK, foobar)
}

func DeleteFoobarByParamId(c *gin.Context) {
	var foobar models.Foobar
	id := c.Params.ByName("id")
	if err := initializers.DB.First(&foobar, id); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": err.Error.Error(),
		})
		log.Panic("Find Error", err.Error.Error())
		return
	}
	initializers.DB.Delete(&foobar, id)
	c.JSON(http.StatusAccepted, foobar)
}

func ShownFoobarByParamReg(c *gin.Context) {
	var foobar []models.Foobar
	reg, _ := strconv.ParseUint(c.Params.ByName("reg"), 10, 64)
	if err := initializers.DB.Where(&models.Foobar{Registration: reg}).Find(&foobar); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": err.Error.Error(),
		})
		log.Panic("Find Error", err.Error.Error())
		return
	}
	c.JSON(http.StatusOK, foobar)
}
