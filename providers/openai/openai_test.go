package openai

import "testing"

func TestCreateOpenAIProvider(t *testing.T) {
	openAIProvider := CreateOpenAIProvider()

	if openAIProvider.Id != "openai" {
		t.Errorf("expected openai as Id, received %s", openAIProvider.Id)
	}
}

func TestGetOpenAIModels(t *testing.T) {
	openAIProvider := CreateOpenAIProvider()

	models, err := GetOpenAIModels(&openAIProvider)
	if err != nil {
		t.Errorf("expected no error, received %v", err)
	}
	if len(models) == 0 {
		t.Errorf("expected model list, received none")
	}

	badModelList := true
	expectedModel := "gpt-4o"
	for _, model := range models {
		if model == expectedModel {
			badModelList = false
		}
	}
	if badModelList {
		t.Errorf("expected model list was bad, received %v", models)
	}

}
