package parser

import (
	"bufio"
	"fmt"
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
	// Split the content into lines
	lines := strings.Split(content, "\n")

	// Initialize a variable to indicate whether we are inside the inputs block
	inputsBlockStarts := -1

	// Initialize a variable to hold the line number of the last closing curly brace
	inputsBlockEnds := -1

	// Iterate through each line of the content
	for lineNumber, line := range lines {

		// Look for the inputs block
		if strings.TrimSpace(strings.Split(line, "=")[0]) == "inputs" {
			// +1 cuz the line number starts from 0
			inputsBlockStarts = lineNumber + 1
			continue
		}

		// If inside the inputs block and the current line contains a closing curly brace, record its line number
		if inputsBlockStarts != -1 && strings.TrimSpace(line) == "}" {
			// +1 cuz the line number starts from 0
			inputsBlockEnds = lineNumber + 1
		}
	}

	log.Printf("Input blocks start at line %d and end at line %d\n", inputsBlockStarts, inputsBlockEnds)

	// If we found the line number of the last closing curly brace, use bufio.Scanner to parse the content above it
	scanner := bufio.NewScanner(strings.NewReader(strings.Join(lines[inputsBlockStarts:inputsBlockEnds-1], "\n")))
	for scanner.Scan() {
		line := scanner.Text()
		// Process the content above the last closing curly brace
		fmt.Println(line)
	}

}

// Skip empty lines and comments
// if len(line) == 0 || strings.Contains(line, "#") {
// 	continue
// }
