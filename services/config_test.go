// services/config_test.go
package services

import (
	"os"
	"testing"

	"github.com/burritocatai/llamacat/providers"
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
