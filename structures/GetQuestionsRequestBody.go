package structures

type GetQuestionsRequestBody struct {
	Questions []QuestionAnswerPair `json:"questions"`
}
