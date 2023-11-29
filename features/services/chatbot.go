package services

import (
	"bufio"
	"context"
	"fmt"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"strings"

	"github.com/sashabaranov/go-openai"
)

type IChatbotService interface {
	Chatbot(client *openai.Client, input request.Chatbot) (response.Chatbot, error)
}

type ChatbotService struct {
	client *openai.Client
}

func NewChatbotService(client *openai.Client) *ChatbotService {
	return &ChatbotService{client}
}

func (cs *ChatbotService) Chatbot(client *openai.Client, input request.Chatbot) {
	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "Response as a customer service operating in the field of local art and cultural event",
			},
		},
		MaxTokens: 150,
	}

	s := bufio.NewScanner(strings.NewReader(input.Message))
	for s.Scan() {
		req.Messages = append(req.Messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: s.Text(),
		})
		resp, err := client.CreateChatCompletion(context.Background(), req)
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			continue
		}
		fmt.Printf("%s\n\n", resp.Choices[0].Message.Content)
		req.Messages = append(req.Messages, resp.Choices[0].Message)
		fmt.Print("> ")
	}
}
