package storage

import (
	"slices"
	"testing"
)

func TestGetObsidianVaults(t *testing.T) {

	expectedVaults := []VaultInfo{
		{Name: "test", Path: "/home/vscode/vault"},
	}

	receivedVaults := GetObsidianVaults()

	if !slices.Contains(receivedVaults, expectedVaults[0]) {
		t.Errorf("vault %s was not found as expected. found %v", expectedVaults[0].Name, receivedVaults)
	}

}
