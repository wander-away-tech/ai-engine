package ai

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func GenerateItinerary(prompt string) (*string, error) {
	llm, err := ollama.New(ollama.WithModel("gemma2:2b"), ollama.WithServerURL("http://ollama:11434"))
	if err != nil {
		return nil, err
	}

	// Run the llm in the background context
	ctx := context.Background()

	response, err := llm.Call(ctx, prompt)

	llm.GenerateContent(ctx, []llms.MessageContent{{
		Role:  llms.ChatMessageTypeSystem,
		Parts: []llms.ContentPart{llms.TextContent{Text: prompt}},
	}})

	fmt.Println(response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
