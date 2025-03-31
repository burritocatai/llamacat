package fake

import (
	"context"

	"github.com/burritocatai/llamacat/providers"
	fake_llm "github.com/tmc/langchaingo/llms/fake"
	"github.com/tmc/langchaingo/prompts"
)

func CreateFakeAIProvider() providers.AIProvider {
	fakeAIProvider := providers.NewAIProvider(
		"FAKEAI_API_KEY",
		"sk-",
		"https://api.fake.com/v1",
		"models",
		"Fake",
		"fakeai",
		"https://api.fake.com/v1",
	)
	fakeAIProvider.Call = func(model string, prompt prompts.PromptTemplate,
		content string, context context.Context) (string, error) {
		return GetFakeAIResponse(fakeAIProvider, model, prompt, content, context)
	}

	// fakeAIProvider.GetModels = func() ([]string, error) {
	// 	return GetOpenAIModels(fakeAIProvider)
	// }

	return *fakeAIProvider
}

func GetFakeAIResponse(provider *providers.AIProvider, model string, prompt prompts.PromptTemplate,
	content string, context context.Context) (string, error) {

	responses := setupResponses()
	fakeLLM := fake_llm.NewFakeLLM(responses)

	result, err := prompt.Format(map[string]any{"content": content})

	completion, err := fakeLLM.Call(context, result)
	return completion, err

}

func setupResponses() []string {
	return []string{
		"I'm a fake AI, but helpful",
	}
}
