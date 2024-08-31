package main

import (
	"encoding/json"
	"net/http"

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

func sendResult(ctx *gin.Context, prompt string) {
	result := ""

	result = strings.Replace(result, "```json", "", -1)
	result = strings.Replace(result, "```", "", -1)
	result = strings.TrimSpace(result)

	bytes, err := json.Marshal(result)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error! The LLM response was not formatted as json!",
			"error":   "ERR_LLM_RESPONSE_NOT_JSON_ONLY",
		})
	}

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

		sendResult(ctx, prompt)
	} else {
		prompt, err := prompts.GetFollowUpQuestionsGeneratePrompt(body.Questions)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error! The JSON body did not have previous questions!",
				"error":   "ERR_JSON_QUESTIONS_NOT_FOUND",
			})
		}

		sendResult(ctx, prompt)
	}
}

func HandleGetItinerary(ctx *gin.Context) {
	body := new(structures.ItineraryRequestBody)
	err := utils.GetJSONBody(ctx, body)
	if err != nil {
		return
	}

	prompt := prompts.GetItineraryGeneratePrompt(body.Destination, body.Duration, body.Preferences)

	sendResult(ctx, prompt)
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
