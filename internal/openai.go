package internal

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

// OpenAIClient represents an OpenAI API client
type OpenAIClient struct {
	client *openai.Client
}

// NewOpenAIClient creates a new OpenAI client with the provided API key
func NewOpenAIClient(apiKey string) *OpenAIClient {
	client := openai.NewClient(apiKey)
	return &OpenAIClient{
		client: client,
	}
}

// GenerateResponse sends a request to the OpenAI API and returns the generated text
func (c *OpenAIClient) GenerateResponse(prompt string) (string, error) {
	resp, err := c.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		return "", fmt.Errorf("error generating response: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response choices returned from OpenAI")
	}

	return resp.Choices[0].Message.Content, nil
}