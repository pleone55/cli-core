package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateProject(t *testing.T) {
	projectName := "testproject"
	defer os.RemoveAll(projectName) // Clean up after test

	os.Args = []string{"generate", projectName}

	main()
	if _, err := os.Stat(projectName); os.IsNotExist(err) {
		t.Fatalf("Expected project directory %s to be created, but it does not exist", projectName)
	}
	if _, err := os.Stat(filepath.Join(projectName, "main.go")); os.IsNotExist(err) {
		t.Fatalf("Expected main.go file to be created in %s, but it does not exist", projectName)
	}
	if _, err := os.Stat(filepath.Join(projectName, "go.mod")); os.IsNotExist(err) {
		t.Fatalf("Expected go.mod file to be created in %s, but it does not exist", projectName)
	}
}
