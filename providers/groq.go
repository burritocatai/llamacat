// internal/providers/groq.go
package providers

import (
	"context"

	"github.com/burritocatai/llamacat/services"
	"github.com/tmc/langchaingo/prompts"
)

// Groq is an OpenAI compatible service. Just need to register it
func init() {
	groqAIProvider := services.NewAIProvider(
		"GROQ_API_KEY",
		"gsk_",
		"https://api.groq.com/openai/v1",
		"models",
		"Groq",
		"groq",
		"https://groq.com",
	)
	groqAIProvider.Call = func(model string, prompt prompts.PromptTemplate,
		content string, context context.Context) (string, error) {
		return GetOpenAIResponse(groqAIProvider, model, prompt, content, context)
	}

	groqAIProvider.GetModels = func() ([]string, error) {
		return GetOpenAIModels(groqAIProvider)
	}

	services.RegisterAIProvider(*groqAIProvider)
}
