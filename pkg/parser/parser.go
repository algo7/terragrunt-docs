package parser

import (
	"bufio"
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
	extractInputsContent(string(content))

	// if inputsContent != "" {
	// 	return inputsContent
	// }

	return "Default Settings"
}

// KeyValue represents a key-value pair.
type KeyValue struct {
	Key   string
	Value string
}

// extractInputsContent extracts the content inside the inputs block from a terragrunt file, accounting for comments
func extractInputsContent(content string) {
	// Create a scanner to read lines from the content
	scanner := bufio.NewScanner(strings.NewReader(content))

	// Initialize a variable to inidicate whether we are inside the inputs block
	insideInputsBlock := false

	// Iterate through each line of the content
	for scanner.Scan() {
		line := scanner.Text()

		// Skip empty lines and comments
		if len(line) == 0 || strings.Contains(line, "#") {
			continue
		}

		// Look for the inputs block
		if strings.TrimSpace(strings.Split(line, "=")[0]) == "inputs" {
			insideInputsBlock = true
			continue
		}

		if insideInputsBlock {
			log.Println(line)
		}

	}
}
