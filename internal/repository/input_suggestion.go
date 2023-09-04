package repository

import (
	"context"
	"github.com/sashabaranov/go-openai"
)

type InputSuggestionQuery interface {
	GeneratePrompt(prompt string) (*string, error)
}

type inputSuggestionQuery struct {
	openAI *openai.Client
}

func NewInputSuggestionQuery(openAI *openai.Client) InputSuggestionQuery {
	return &inputSuggestionQuery{
		openAI: openAI,
	}
}

func (i *inputSuggestionQuery) GeneratePrompt(prompt string) (*string, error) {
	resp, err := i.openAI.CreateChatCompletion(
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
		return nil, err
	}

	FirstChoice := resp.Choices[0].Message.Content
	return &FirstChoice, nil
}
