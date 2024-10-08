package assistant

import "context"

type openAIAssistant struct{}

func NewOpenAIAssistant() *openAIAssistant {
	return &openAIAssistant{}
}

func (ai *openAIAssistant) GenerateMessage(ctx context.Context, diffInput string) (string, error) {
	return "", nil
}
