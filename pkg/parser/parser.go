package parser

import (
	"bufio"
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
		return inputsContent
	}

	return "Default Settings"
}

func extractInputsContent(content string) string {
	// Create a scanner to read lines from the content
	scanner := bufio.NewScanner(strings.NewReader(content))

	// Initialize a flag to indicate whether we are inside the inputs block
	insideInputsBlock := false

	// Create a buffer to store the extracted content
	var result strings.Builder

	// Iterate through each line of the content
	for scanner.Scan() {
		line := scanner.Text()

		// Check if we are inside the inputs block
		if insideInputsBlock {
			// If the line is empty or contains '#', it's a comment; skip it
			if len(line) == 0 || strings.Contains(line, "#") {
				continue
			}

			// If the line starts with '}', we've reached the end of the inputs block
			if strings.HasPrefix(line, "}") {
				break
			}

			// Append the line to the result
			result.WriteString(line)
			result.WriteString("\n")
		} else {
			// Check if the line starts with "inputs = {"
			if strings.HasPrefix(line, "inputs = {") {
				insideInputsBlock = true
			}
		}
	}

	// Return the content inside the curly braces as a string
	return result.String()
}
