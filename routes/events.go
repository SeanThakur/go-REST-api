package routes

import (
	"net/http"
	"seanThakur/go-restapi/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LooseDict map[string]any

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, LooseDict{"message": "Could not fetch event."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadGateway, LooseDict{"message": "could not find event"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, LooseDict{"message": "server error"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, LooseDict{"message": "Could not parse request data."})
		return
	}
	event.Id = 1
	event.UserId = 1
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, LooseDict{"message": "error while creating event"})
	}
	context.JSON(http.StatusCreated, LooseDict{"message": "Event Created", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadGateway, LooseDict{"message": "could not find parse event id"})
		return
	}

	_, err = models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, LooseDict{"message": "Could not find event id"})
		return
	}

	var updateEvent models.Event
	err = context.ShouldBindJSON(&updateEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, LooseDict{"message": "Could not parse data"})
		return
	}

	updateEvent.Id = eventId
	err = updateEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, LooseDict{"message": "Could not update event, please try again in sometime."})
	}

	context.JSON(http.StatusOK, LooseDict{"message": "Event updated successfully!"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadGateway, LooseDict{"message": "could not find parse event id"})
		return
	}

	result, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, LooseDict{"message": "Could not find event id"})
		return
	}

	err = result.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, LooseDict{"message": "Could not delete this event"})
		return
	}

	context.JSON(http.StatusOK, LooseDict{"message": "Event deleted successfully"})
}
