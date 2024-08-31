package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gomods.euniz.com/gomods/ai-engine/ai"
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

	result, err := ai.GenerateItinerary(prompt)

	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error! The LLM response was not formatted as json!",
			"error":   "ERR_LLM_RESPONSE_NOT_JSON_ONLY",
		})
	}

	*result = strings.Replace(*result, "```json", "", -1)
	*result = strings.Replace(*result, "```", "", -1)
	*result = strings.TrimSpace(*result)

	bytes, err := json.Marshal(*result)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error! The LLM response was not formatted as json!",
			"error":   "ERR_LLM_RESPONSE_NOT_JSON_ONLY",
		})
	}

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
