package assistant

import (
	"context"
	"net/http"
	"net/url"

	"github.com/ollama/ollama/api"
)

type ollamaAssistant struct {
	model string
}

func NewOllamaAssistant(cfg Config) *ollamaAssistant {
	return &ollamaAssistant{
		model: cfg.Model,
	}
}

func (ai *ollamaAssistant) GenerateMessage(ctx context.Context, diffInput string) (string, error) {
	url, err := url.Parse("http://localhost:11434")
	if err != nil {
		return "", err
	}

	httpClient := http.Client{}
	client := api.NewClient(url, &httpClient)

	messages := []api.Message{
		{
			Role:    "system",
			Content: systemPrompt,
		},
		{
			Role:    "user",
			Content: userPrompt,
		},
		{
			Role:    "assistant",
			Content: assistantPrompt,
		},
		{
			Role:    "user",
			Content: diffInput,
		},
	}

	req := &api.ChatRequest{
		Model:    ai.model,
		Messages: messages,
		Stream:   new(bool),
	}

	var resp string
	err = client.Chat(ctx, req, func(cr api.ChatResponse) error {
		resp += cr.Message.Content

		return nil
	})
	if err != nil {
		return "", err
	}

	return resp, nil
}
