// internal/providers/groq.go
package groq

import (
	"context"
	"fmt"

	"github.com/burritocatai/llamacat/providers"
	"github.com/burritocatai/llamacat/providers/openai"
	"github.com/tmc/langchaingo/prompts"
)

type GroqModel struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

type GroqModels struct {
	Models []GroqModel `json:"models"`
}

var GroqModelList = GroqModels{
	Models: []GroqModel{
		{
			ID:          "o3-mini",
			Description: "fast, reasoning",
		},
		{
			ID:          "o1-mini",
			Description: "smaller, fast reasoning",
		},
		{
			ID:          "o1",
			Description: "high intelligence reasoning model",
		},
		{
			ID:          "o1-pro",
			Description: "o1 with more compute",
		},
		{
			ID:          "gpt-4.5-preview",
			Description: "largest and most capable gpt model",
		},
		{
			ID:          "gpt-4o",
			Description: "fast intelligence flexible gpt model",
		},
		{
			ID:          "chatgpt-4o",
			Description: "model used in chat gpt",
		},
		{
			ID:          "gpt-4o-mini",
			Description: "fast affordable small model",
		},
	},
}

func GetGroqModels(provider *providers.AIProvider) ([]string, error) {
	modelList := make([]string, 0)
	for _, model := range GroqModelList.Models {
		modelList = append(modelList, model.ID)
	}
	if len(modelList) == 0 {
		return modelList, fmt.Errorf("no models returned for %s", provider.Id)
	}

	return modelList, nil
}

func CreateGroqProvider() providers.AIProvider {
	groqAIProvider := providers.NewAIProvider(
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
		return openai.GetOpenAIResponse(groqAIProvider, model, prompt, content, context)
	}

	groqAIProvider.GetModels = func() ([]string, error) {
		return openai.GetOpenAIModels(groqAIProvider)
	}

	return *groqAIProvider
}

// Groq is an OpenAI compatible service. Just need to register it
func init() {
	providers.RegisterAIProvider(CreateGroqProvider())
}
