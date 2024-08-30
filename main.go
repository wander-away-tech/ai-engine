package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ItineraryRequestBody struct {
	Destination string `json:destination`
	Duration    string `json:duration`
	Preferences string `json:preferences`
}

func main() {
	r := gin.Default()
	api := r.Group("/api")
	api.POST("/itinerary", func(ctx *gin.Context) {
		bodyBytes, err := io.ReadAll(ctx.Request.Body)
		body := new(ItineraryRequestBody)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad Request, Unable to read Request Body",
				"error":   "ERR_BAD_REQUEST_BODY",
			})
			return
		}
		http.Get()
		json.Unmarshal(bodyBytes, &body)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Lol",
		})
	})
	r.Run(":5000") // listen and serve on 0.0.0.0:8080
}
