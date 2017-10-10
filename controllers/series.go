package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ileyd/topaz/models"
)

//SeriesController ...
type SeriesController struct{}

var seriesModel = new(models.SeriesModel)

//All ...
func (ctrl SeriesController) All(c *gin.Context) {
	var series, err = seriesModel.GetAll()
	if err != nil {
		c.JSON(406, gin.H{"message": "Series could not be retrieved", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, series)
}

//One ...
func (ctrl SeriesController) One(c *gin.Context) {
	var kitsuID = c.Param("kitsuID")
	kitsuIDInt, err := strconv.ParseInt(kitsuID, 10, 64)
	if err != nil {
		c.JSON(406, gin.H{"message": "Kitsu ID could not be parsed", "error": err.Error()})
		c.Abort()
		return
	}
	series, err := seriesModel.GetOne("kitsuID", kitsuIDInt)
	if err != nil {
		c.JSON(406, gin.H{"message": "Series could not be retrieved", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, series)
}
