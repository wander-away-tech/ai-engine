package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gomods.euniz.com/gomods/ai-engine/prompts"
	"gomods.euniz.com/gomods/ai-engine/utils"
)

type ItineraryRequestBody struct {
	Destination string `json:"destination"`
	Duration    string `json:"duration"`
	Preferences string `json:"preferences"`
}

func HandleGetQuestion(ctx *gin.Context) {
}

func HandleGetItinerary(ctx *gin.Context) {
	body := new(ItineraryRequestBody)
	err := utils.GetJSONBody(ctx, body)
	if err != nil {
		return
	}

	prompt := prompts.GetItineraryGeneratePrompt(body.Destination, body.Duration, body.Preferences)

	result := "```json {\"name\": \"Alice\", \"age\": 30, \"city\": \"New York\", \"occupation\": \"Software Engineer\", \"is_married\": true} ```"

	result = strings.Replace(result, "```json", "", -1)
	result = strings.Replace(result, "```", "", -1)
	result = strings.TrimSpace(result)

	bytes, err := json.Marshal(result)

	ctx.Data(http.StatusOK, "application/json", bytes)
}

func main() {
	// The Main Gin Server
	server := gin.Default()

	// The API Route Group
	api := server.Group("/api")

	api.POST("/itinerary", HandleGetItinerary)
	api.POST("/questions", HandleGetQuestion)

	server.Run(":5000") // listen and serve on 0.0.0.0:5000
}
