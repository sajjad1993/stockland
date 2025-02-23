package openai

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

const Model = openai.GPT3Dot5Turbo

type Processor struct {
	apiKey string
	client *openai.Client
}

func NewOpenAIProcessor(apiKey string) *Processor {

	client :=
		new(openai.Client)
	return &Processor{client: client,
		apiKey: apiKey,
	}
}

func (p *Processor) ProcessRequest(prompt string) (string, error) {

	client := openai.NewClient(p.apiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: Model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
