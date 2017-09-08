package controllers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ileyd/topaz/models"
	"github.com/ileyd/topaz/utils/handlers"
)

//SonarrEventController ...
type SonarrEventController struct{}

var eventModel = new(models.SonarrEventModel)

//Create ...
func (ctrl SonarrEventController) Create(c *gin.Context) {
	var event models.SonarrEvent

	if err := c.BindJSON(&event); err != nil {
		c.JSON(406, gin.H{"message": "Invalid event", "form": event})
		log.Println("Bind failed SEC.Create", err)
		log.Println(event)
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Event received"})

	handlers.HandleSonarrEventRegistration(event)
}

//All ...
func (ctrl SonarrEventController) All(c *gin.Context) {
	var events = eventModel.GetAll()
	c.JSON(200, events)
}

//One ...
func (ctrl SonarrEventController) One(c *gin.Context) {
	var id = c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(406, gin.H{"message": "Event ID could not be parsed", "error": err.Error()})
		c.Abort()
		return
	}
	var event = eventModel.Get(idInt)
	c.JSON(200, event)
}
