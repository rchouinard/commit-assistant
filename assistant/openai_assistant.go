package assistant

import "context"

type openAIAssistant struct{}

func NewOpenAIAssistant() Assistant {
	return &openAIAssistant{}
}

func (ai *openAIAssistant) GenerateMessage(ctx context.Context, msgs []string) (string, error) {
	return "", nil
}
