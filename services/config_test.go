// services/config_test.go
package services

import (
	"os"
	"testing"

	"github.com/burritocatai/llamacat/providers"
	"github.com/burritocatai/llamacat/providers/fake"
	"github.com/tmc/langchaingo/prompts"
)

func TestGetAPIKey(t *testing.T) {
	testAPIKey := "catsAreAPIs"
	os.Setenv("TEST_API_KEY", testAPIKey)

	testProvider := providers.NewAIProvider("TEST_API_KEY", "", "", "", "", "", "")
	apikey, err := GetAPIKey(testProvider)
	expected := testAPIKey

	if apikey != expected || err != nil {
		t.Errorf("expected '%s' but got '%s'", expected, apikey)
	}
}

func TestGetProviderAndModel(t *testing.T) {

	inputModel := "fakeai:fake-model-9"
	fakeAIProvider := fake.CreateFakeAIProvider()
	providers.RegisterAIProvider(fakeAIProvider)
	expectedModel := "fake-model-9"

	receivedProvider, receivedModel, err := GetProviderAndModel(inputModel)

	if receivedProvider.Id != fakeAIProvider.Id {
		t.Errorf("expected to receive provider %s, received %s", fakeAIProvider.Id, receivedProvider.Id)
	}
	if receivedModel != expectedModel {
		t.Errorf("expected to receive model %s, received %s", expectedModel, receivedModel)
	}

	if err != nil {
		t.Errorf("did not expect error. received %v", err)
	}
}

func TestGetPrompt(t *testing.T) {
	inputPrompt := "fake:prompt"
	expectedPrompt := prompts.NewPromptTemplate("You are a helpful assistant. Help the user with their content.\n\nCONTENT: {{.content}}", []string{"content"})

	receivedPrompt, err := GetPrompt(inputPrompt)
	if err != nil {
		t.Errorf("did not expect error. received %v", err)
	}

	expectedPromptFormatted, _ := expectedPrompt.FormatPrompt(map[string]any{"content": "test"})
	receivedPromptFormatted, err := receivedPrompt.FormatPrompt(map[string]any{"content": "test"})
	if err != nil {
		t.Errorf("did not expect error. received %v", err)
	}

	if expectedPromptFormatted != receivedPromptFormatted {
		t.Errorf("did not expect to receive this prompt, received %v, expected %v", receivedPromptFormatted, expectedPromptFormatted)
	}

	if err != nil {
		t.Errorf("did not expect error. received %v", err)
	}

}
