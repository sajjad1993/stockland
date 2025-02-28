package openai

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"stockland/service/product/domain"
)

const Model = openai.GPT4oMini

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

func (p *Processor) ProcessRequest(ctx context.Context, prompt string, sub string) (domain.Offer, error) {
	client := openai.NewClient(p.apiKey)
	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    Model,
			Messages: getPrompt(prompt, sub),
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return domain.Offer{}, err
	}
	offer := domain.NewOffer(resp.Choices[0].Message.Content)
	return offer, nil
}

func getPrompt(prompt string, sub string) []openai.ChatCompletionMessage {
	return []openai.ChatCompletionMessage{
		{
			Role: openai.ChatMessageRoleSystem,
			Content: fmt.Sprintf(
				"You are a sales manager specializing in stock car tools. "+
					"Never send links to users. "+
					"Only suggest from these products: %s. "+
					"Help the user choose which one works best for them.",
				sub,
			),
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: prompt,
		},
	}
}

func (p *Processor) AnalyzeImage(ctx context.Context, imagePath string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}
