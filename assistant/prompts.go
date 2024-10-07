package assistant

import _ "embed"

//go:embed prompts/system.md
var systemPrompt string

//go:embed prompts/assistant.md
var assistantPrompt string

//go:embed prompts/user.md
var userPrompt string
