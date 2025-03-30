// services/config_test.go
package services

import (
	"os"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	testAPIKey := "catsAreAPIs"
	os.Setenv("TEST_API_KEY", testAPIKey)

	testProvider := NewAIProvider("TEST_API_KEY", "", "", "", "", "", "")
	apikey, err := GetAPIKey(testProvider)
	expected := testAPIKey

	if apikey != expected || err != nil {
		t.Errorf("expected '%s' but got '%s'", expected, apikey)
	}
}

func TestGetProviderAndModel(t *testing.T) {
	testModelFlag := "openai:gpt-4o-mini"
	testProvider := NewAIProvider(
		"OPENAI_API_KEY",
		"sk-",
		"https://api.openai.com/v1",
		"models",
		"OpenAI",
		"openai",
		"https://platform.openai.com",
	)
	RegisterAIProvider(*testProvider)

	expectedProviderId, expectedModel := "openai", "gpt-4o-mini"

	provider, model, err := GetProviderAndModel(testModelFlag)

	if provider.Id != expectedProviderId {
		t.Errorf("expected '%s' but got '%s'", expectedProviderId, provider.Id)
	}

	if model != expectedModel {
		t.Errorf("expected '%s' but got '%s'", expectedModel, model)
	}

	if err != nil {
		t.Errorf("expected no error but got %v", err)
	}

}
