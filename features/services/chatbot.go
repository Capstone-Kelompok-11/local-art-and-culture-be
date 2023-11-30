package services

import (
	"bufio"
	"context"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"strings"

	"github.com/sashabaranov/go-openai"
)

type IChatbotService interface {
	Chatbot(client openai.Client, input request.Chatbot) (response.Chatbot, error)
}

type ChatbotService struct {
	client *openai.Client
}

func NewChatbotService(client *openai.Client) *ChatbotService {
	return &ChatbotService{client}
}

func (cs *ChatbotService) Chatbot(client openai.Client, input request.Chatbot) (response.Chatbot, error) {
	cs.client = &client
	var res response.Chatbot
	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "Response as a customer service operating in the field of local art and cultural events, originating from a company named Lokasani.",
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
		resp, err := cs.client.CreateChatCompletion(context.Background(), req)
		if err != nil {
			continue
		}
		res.Message = resp.Choices[0].Message.Content
		req.Messages = append(req.Messages, resp.Choices[0].Message)
	}
	return res, nil
}
