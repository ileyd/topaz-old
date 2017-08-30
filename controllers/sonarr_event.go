package controllers

import (
	"strconv"

	"github.com/ileyd/topaz/forms"
	"github.com/ileyd/topaz/models"

	"github.com/gin-gonic/gin"
)

//SonarrEventController ...
type SonarrEventController struct{}

var eventModel = new(models.SonarrEvent)

//Create ...
func (ctrl SonarrEventController) Create(c *gin.Context) {
	userID := getUserID(c)

	if userID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	var eventForm forms.SonarrEventForm

	if c.BindJSON(&eventForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": eventForm})
		c.Abort()
		return
	}

	eventID, err := eventModel.Create(userID, eventForm)

	if eventID > 0 && err != nil {
		c.JSON(406, gin.H{"message": "Event could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Event created", "id": eventID})
}

//All ...
func (ctrl SonarrEventController) All(c *gin.Context) {
	userID := getUserID(c)

	if userID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	data, err := eventModel.All(userID)

	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the events", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": data})
}

//One ...
func (ctrl SonarrEventController) One(c *gin.Context) {
	userID := getUserID(c)

	if userID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	id := c.Param("id")

	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

		data, err := eventModel.One(userID, id)
		if err != nil {
			c.JSON(404, gin.H{"Message": "Event not found", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"data": data})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter"})
	}
}

//Update ...
func (ctrl SonarrEventController) Update(c *gin.Context) {
	userID := getUserID(c)

	if userID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	id := c.Param("id")
	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

		var eventForm forms.EventForm

		if c.BindJSON(&eventForm) != nil {
			c.JSON(406, gin.H{"message": "Invalid parameters", "form": eventForm})
			c.Abort()
			return
		}

		err := eventModel.Update(userID, id, eventForm)
		if err != nil {
			c.JSON(406, gin.H{"Message": "Event could not be updated", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"message": "Event updated"})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter", "error": err.Error()})
	}
}

//Delete ...
func (ctrl SonarrEventController) Delete(c *gin.Context) {
	userID := getUserID(c)

	if userID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	id := c.Param("id")
	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

		err := eventModel.Delete(userID, id)
		if err != nil {
			c.JSON(406, gin.H{"Message": "Event could not be deleted", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"message": "Event deleted"})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter"})
	}
}
