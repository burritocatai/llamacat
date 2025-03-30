// internal/providers/mistral.go
package providers

import (
	"context"

	"github.com/tmc/langchaingo/prompts"
)

func init() {
	mistralProvider := NewAIProvider(
		"",
		"MISTRAL_API_KEY",
		"xx",
		"https://api.mistral.ai/v1",
		"models",
		"Mistral",
		"mistral",
		"https://console.mistral.ai/home",
	)
	mistralProvider.Call = func(model string, prompt prompts.PromptTemplate,
		content string, context context.Context) (string, error) {
		return GetOpenAIResponse(mistralProvider, model, prompt, content, context)
	}

	// mistral API is similar enough to use the OpenAI model pull
	mistralProvider.GetModels = func() ([]string, error) {
		return GetOpenAIModels(mistralProvider)
	}

	RegisterAIProvider(*mistralProvider)
}
