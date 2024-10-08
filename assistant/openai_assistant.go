package assistant

import (
	"context"
	"os"

	"github.com/openai/openai-go"
)

type openAIAssistant struct {
	apiKey  string
	baseUrL string
	model   string
}

func NewOpenAIAssistant(cfg Config) *openAIAssistant {
	if cfg.APIKey == "" {
		cfg.APIKey = os.Getenv("OPENAI_API_KEY")
	}

	if cfg.BaseURL == "" {
		cfg.BaseURL = "https://api.openai.com/v1/"
	}

	if cfg.Model == "" {
		cfg.BaseURL = openai.ChatModelGPT3_5Turbo
	}

	return &openAIAssistant{
		apiKey:  cfg.APIKey,
		baseUrL: cfg.BaseURL,
		model:   cfg.Model,
	}
}

func (ai *openAIAssistant) GenerateMessage(ctx context.Context, diffInput string) (string, error) {
	return "", nil
}
