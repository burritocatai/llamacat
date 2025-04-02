package prompts

import (
	"fmt"
	"testing"
)

func TestDownloadDefaultPrompts(t *testing.T) {

	status, err := DownloadDefaultPrompts()

	if err != nil {
		t.Errorf("received err did not expect %v", err)
	}
	if status == Cloned || status == AlreadyExists {
		// No error
	} else {
		t.Errorf("received bad status %v when %v or %v expected", status, Cloned, AlreadyExists)
	}

}

func TestUpdateDefaultPrompts(t *testing.T) {

	status, err := UpdateDefaultPrompts()

	if err != nil {
		t.Errorf("received err did not expect %v", err)
	}
	if status == UpToDate || status == Updated {
		// No error
	} else {
		t.Errorf("received bad status %v when %v or %v expected", status, UpToDate, Updated)
	}
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
