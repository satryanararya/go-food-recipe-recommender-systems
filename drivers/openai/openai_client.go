package openai

import (
	"context"
	"fmt"

	err_util "github.com/satryanararya/go-chefbot/utils/error"
	oai "github.com/sashabaranov/go-openai"
)

type OpenAIClient interface {
	GetRecommendation(prompt string, histories []string) (string, error)
}

type openAIClient struct {
	client *oai.Client
}

func NewOpenAIClient(apiKey string) *openAIClient {
	return &openAIClient{
		client: oai.NewClient(apiKey),
	}
}

func (o *openAIClient) GetRecommendation(prompt string, histories []string) (string, error) {
	ctx := context.Background()

	messages := []oai.ChatCompletionMessage{
		{
			Role: oai.ChatMessageRoleSystem,
			Content: prompt,
		},
	}
	if len(histories) > 0 {
		for _, v := range histories {
			messages = append(messages, oai.ChatCompletionMessage{
				Role:    oai.ChatMessageRoleAssistant,
				Content: v,
			})
		}
	}
	req := oai.ChatCompletionRequest{
		Model:    oai.GPT3Dot5Turbo,
		Messages: messages,
	}
	res, err := o.client.CreateChatCompletion(ctx, req)
	if err != nil {
		fmt.Println(err)
		return "", err_util.ErrExternalService
	}

	return res.Choices[0].Message.Content, nil
}