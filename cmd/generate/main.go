package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: generate [command]")
		os.Exit(1)
	}
	projectName := os.Args[1]

	err := os.Mkdir(projectName, 0755)
	if err != nil {
		fmt.Println("Error creating project directory:", err)
		os.Exit(1)
	}

	// Create main.go file from the mainTemplate.txt file in the template directory
	templatePath := filepath.Join("../templates", "mainTemplate.txt")
	templateContent, err := os.ReadFile(templatePath)
	if err != nil {
		fmt.Println("Error reading main.go template file:", err)
		os.Exit(1)
	}
	// Create the main.go file in the project directory
	mainFilePath := filepath.Join(projectName, "main.go")
	err = os.WriteFile(mainFilePath, []byte(templateContent), 0644)
	if err != nil {
		fmt.Println("Error creating main.go file:", err)
		os.Exit(1)
	}

	// Create go.mod file from the goModTemplate.txt file in the template directory
	goModTemplatePath := filepath.Join("../templates", "modTemplate.txt")
	goModTemplateContent, err := os.ReadFile(goModTemplatePath)
	if err != nil {
		fmt.Println("Error reading go.mod template file:", err)
		os.Exit(1)
	}
	goModFilePath := filepath.Join(projectName, "go.mod")
	err = os.WriteFile(goModFilePath, goModTemplateContent, 0644)
	if err != nil {
		fmt.Println("Error creating go.mod file:", err)
		os.Exit(1)
	}
	fmt.Println("Project created successfully at:", projectName)
	flag.Parse()
	fmt.Println("Project structure:")
	for _, file := range []string{"main.go", "go.mod"} {
		fmt.Println(" -", filepath.Join(projectName, file))
	}
}
