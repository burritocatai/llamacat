package prompts

import (
	"fmt"
	"testing"
)

func TestDownloadDefaultPrompts(t *testing.T) {

	err := DownloadDefaultPrompts()

	if err != nil {
		t.Errorf("received err did not expect %v", err)
	}
	t.Errorf("test not implemented")

}

func TestUpdateDefaultPrompts(t *testing.T) {

	err := UpdateDefaultPrompts()

	if err != nil {
		t.Errorf("received err did not expect %v", err)
	}
	t.Errorf("test not implemented")

}

func TestDownloadPromptRepo(t *testing.T) {

}

func TestUpdatePromptRepo(t *testing.T) {

}

func TestGetAvailablePrompts(t *testing.T) {
	prompts, err := GetAvailablePrompts("default")
	fmt.Printf("avaialbe prompts are: %v", prompts)

	if err != nil {
		t.Errorf("did not expect error. received %v", err)
	}
}
