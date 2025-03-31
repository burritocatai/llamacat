package groq

import "testing"

func TestCreateGroqProvider(t *testing.T) {
	groqProvider := CreateGroqProvider()

	if groqProvider.Id != "groq" {
		t.Errorf("expected groq as Id, received %s", groqProvider.Id)
	}
}
