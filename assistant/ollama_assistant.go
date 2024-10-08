package assistant

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/ollama/ollama/api"
)

type ollamaAssistant struct {
	baseURL string
	model   string
}

func NewOllamaAssistant(cfg Config) *ollamaAssistant {
	if cfg.BaseURL == "" {
		cfg.BaseURL = os.Getenv("OLLAMA_HOST")
	}

	if cfg.BaseURL == "" {
		cfg.BaseURL = "http://localhost:11434"
	}

	if cfg.Model == "" {
		cfg.Model = "mistral"
	}

	return &ollamaAssistant{
		baseURL: cfg.BaseURL,
		model:   cfg.Model,
	}
}

func (ai *ollamaAssistant) GenerateMessage(ctx context.Context, diffInput string) (string, error) {
	url, err := url.Parse(ai.baseURL)
	if err != nil {
		return "", err
	}

	httpClient := http.Client{
		Timeout: time.Second * 30,
	}
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
