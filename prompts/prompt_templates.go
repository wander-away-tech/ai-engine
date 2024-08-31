package prompts

import (
	"encoding/json"
	"fmt"

	"gomods.euniz.com/gomods/ai-engine/structures"
)

const GENERATE_PROMPT_BASE = "Please do not give any output other than JSON, please make sure you give only JSON output. Generate an itinerary for a person to travel to %v for a duration of %v, they prefer travelling to places like %v."

func GetItineraryGeneratePrompt(destination string, duration string, preferences string) string {
	return fmt.Sprintf(GENERATE_PROMPT_BASE, destination, duration, preferences)
}

// FIXME: CHANGE LATER
const GENERATE_INITIAL_QUESTIONS_PROMPT_BASE = "Please do not give any output other than JSON, please make sure you give only JSON output. Generate 5 questions to understand the travel preferences of a traveller."

/*func GetInitialQuestionsGeneratePrompt() string {
	return fmt.Sprintf(GENERATE_INITIAL_QUESTIONS_PROMPT_BASE)
}*/

const GENERATE_FOLLOW_UP_QUESTIONS_PROMPT_BASE = "Please do not give any output other than JSON, please make sure you give only JSON output. Generate follow up questions. Here are the previous questions and answers in JSON format. \n%v"

func GetFollowUpQuestionsGeneratePrompt(qapair []structures.QuestionAnswerPair) (string, error) {
	bytes, err := json.Marshal(qapair)
	if err != nil {
		return "",err
	}
	// The : is needed because byte arrays cannot be directly turned to a string while slices can.
	str := string(bytes[:])
	return fmt.Sprintf(GENERATE_FOLLOW_UP_QUESTIONS_PROMPT_BASE, str), nil
}
