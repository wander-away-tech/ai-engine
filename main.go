package main

import (
	"encoding/json"
	//"go/types"
	"strings"

	"github.com/gin-gonic/gin"
	"gomods.euniz.com/gomods/ai-engine/prompts"
	"gomods.euniz.com/gomods/ai-engine/structures"
	"gomods.euniz.com/gomods/ai-engine/utils"
)

func isStructEmpty(s structures.GetQuestionsRequestBody) bool {
	return len(s.Questions) == 0
}

func getResult(ctx *gin.Context, prompt string) {
	result := ""

	result = strings.Replace(result, "```json", "", -1)
	result = strings.Replace(result, "```", "", -1)
	result = strings.TrimSpace(result)

	bytes, err := json.Marshal(result)

	ctx.Data(200, "application/json", bytes)
}

func HandleGetQuestion(ctx *gin.Context) {
	body := new(structures.GetQuestionsRequestBody)
	err := utils.GetJSONBody(ctx, body)
	if err != nil {
		return
	}
	if isStructEmpty(*body) {
		prompt := prompts.GENERATE_INITIAL_QUESTIONS_PROMPT_BASE

		getResult(ctx, prompt)
	} else {
		prompt := prompts.GetFollowUpQuestionsGeneratePrompt()

		getResult(ctx, prompt)
	}
}

func HandleGetItinerary(ctx *gin.Context) {
	body := new(structures.ItineraryRequestBody)
	err := utils.GetJSONBody(ctx, body)
	if err != nil {
		return
	}

	prompt := prompts.GetItineraryGeneratePrompt(body.Destination, body.Duration, body.Preferences)

	getResult(ctx, prompt)
	/*result := "```json {\"name\": \"Alice\", \"age\": 30, \"city\": \"New York\", \"occupation\": \"Software Engineer\", \"is_married\": true} ```"

	result = strings.Replace(result, "```json", "", -1)
	result = strings.Replace(result, "```", "", -1)
	result = strings.TrimSpace(result)

	bytes, err := json.Marshal(result)

	ctx.Data(200, "application/json", bytes)*/

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
