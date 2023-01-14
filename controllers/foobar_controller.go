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

// @BasePath /api/v1
// CreateFoobar godoc
// @Summary Creating Foobar
// @Description Create Foobar in database
// @Tags foobar
// @Accept json
// @Produce json
// @Param foobar body models.Foobar true "Foobar Data"
// @Success 200 {object} models.Foobar
// @Router /foobar [post]
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

// @BasePath /api/v1
// ShownFoobar godoc
// @Summary Showning all Foobar
// @Description Shown all Foobar in database
// @Tags foobar
// @Accept json
// @Produce json
// @Success 200 {object} models.Foobar
// @Router /foobar [get]
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

// @BasePath /api/v1
// ShownFoobarByParamId godoc
// @Summary Showning Foobar by ID
// @Description Shown Foobar by ID in database
// @Tags foobar
// @Accept json
// @Produce json
// @Param id path int true "Foobar ID"
// @Success 200 {object} models.Foobar
// @Router /foobar/{id} [get]
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

// @BasePath /api/v1
// EditFoobarByParamId godoc
// @Summary Editing Foobar by ID
// @Description Edit Foobar by ID in database
// @Tags foobar
// @Accept json
// @Produce json
// @Param id path int true "Foobar ID"
// @Param foobar body models.Foobar true "Foobar Data"
// @Success 200 {object} models.Foobar
// @Router /foobar/{id} [patch]
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

// @BasePath /api/v1
// DeleteFoobarByParamId godoc
// @Summary Deleting Foobar by ID
// @Description Delete Foobar by ID in database
// @Tags foobar
// @Accept json
// @Produce json
// @Param id path int true "Foobar ID"
// @Success 200 {object} models.Foobar
// @Router /foobar/{id} [delete]
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

// @BasePath /api/v1
// ShownFoobarByParamReg godoc
// @Summary Showning Foobar by Registration
// @Description Shown Foobar by Registration in database
// @Tags foobar
// @Accept json
// @Produce json
// @Param reg path uint64 true "Foobar Registration"
// @Success 200 {object} models.Foobar
// @Router /foobar/{reg} [get]
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

// @BasePath /api/v1
// ShownFoobarMockToLearnTests godoc
// @Summary Showning all Foobar Mocks
// @Description Shown all Foobar Mocks
// @Tags foobar
// @Accept json
// @Produce json
// @Success 200 {object} models.Foobar
// @Router /foobar/mock [get]
func ShownFoobarMockToLearnTests(c *gin.Context) {
	//Mock reference
	c.JSON(200, helpers.MyFoobar)
	logs.PopulateErrorLogFile("INFO", nil, "Successful call to endpoint: "+""+c.Request.Method+""+c.Request.RequestURI)
}
