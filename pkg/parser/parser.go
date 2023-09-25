package parser

import (
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/algo7/terragrunt-docs/pkg/utils"
)

const (
	inputsBlockRegex = `(?s)inputs\s*=\s*{(.*)}`
)

var (
	inputsBlockPattern = regexp.MustCompile(inputsBlockRegex)
)

// ExtractInputsFromTerragrunt extracts the content inside the inputs block from a terragrunt file
func ExtractInputsFromTerragrunt(file string) string {
	// Read the file
	content, err := os.ReadFile(file)
	utils.ErrorHandler(err)

	// Use the extractInputsContent function to get the content inside the inputs block
	inputsContent := extractInputsContent(string(content))
	if inputsContent != "" {
		log.Println(inputsContent)
		return inputsContent
	}

	return "Default Settings"
}

// extractInputsContent extracts the content inside the inputs block from a terragrunt file
func extractInputsContent(content string) string {
	// Find the start of the "inputs" keyword
	inputsPos := strings.Index(content, "inputs")
	if inputsPos == -1 {
		return ""
	}

	// Find the first opening curly brace after "inputs"
	startPos := strings.Index(content[inputsPos:], "{")
	if startPos == -1 {
		return ""
	}
	startPos += inputsPos + 6 // Adjust the position relative to the entire content

	// Initialize a counter for open curly braces
	counter := 1 // Start with 1 because we've already found the opening brace

	// Start iterating from the character after the found "{"
	for i := startPos + 1; i < len(content); i++ {
		if content[i] == '{' {
			counter++
		} else if content[i] == '}' {
			counter--
		}

		// If the counter returns to zero, we've found the end of the inputs block
		if counter == 0 {
			// Return the content inside the curly braces
			return content[startPos+1 : i]
		}
	}

	// If we reach here, there's no matching closing brace
	return ""
}
