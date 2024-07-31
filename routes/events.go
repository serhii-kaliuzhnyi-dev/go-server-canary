package routes

import (
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
)

func GetEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Something really bad happened"})
	}

	context.JSON(http.StatusOK, events)
}
