// services/config_test.go
package services

import (
	"os"
	"testing"

	"github.com/burritocatai/llamacat/providers"
	"github.com/burritocatai/llamacat/providers/fake"
	"github.com/spf13/viper"
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

	receivedPrompt, err := GetPromptConfig(inputPrompt)
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

func TestGetOutputFunc(t *testing.T) {
	outputParam := "work:Notes/fromllamacat"
	setupViper(t)

	outputFunc, path, target, err := GetOutputConfig(outputParam)
	if err != nil {
		t.Errorf("did not expect an error. received %v", err)
	}
	if outputFunc == nil {
		t.Errorf("did not expect a nil output fun. received nil")
	} else {
		//
		outputFunc("test", path, target)
	}
}

func setupViper(t *testing.T) {
	viper.SetConfigName("test_config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // Look for config in the current directory
	err := viper.ReadInConfig()
	if err != nil {
		t.Fatalf("Error reading config file: %v", err)
	}
}
