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

	// Initialize a slice to store the extracted key-value pairs
	var keyValuePairs []KeyValue

	// Initialize a stack to keep track of the current nesting level
	var nestingStack []string

	// Iterate through each line of the content
	for lineNumber, line := range lines {

		// Look for the inputs block
		if strings.TrimSpace(strings.Split(line, "=")[0]) == "inputs" {
			// +1 because the line number starts from 0
			inputsBlockStarts = lineNumber + 1
			continue
		}

		// If inside the inputs block and the current line contains a closing curly brace, record its line number
		if inputsBlockStarts != -1 && strings.TrimSpace(line) == "}" {
			// +1 because the line number starts from 0
			inputsBlockEnds = lineNumber + 1
		}
	}

	log.Printf("Input blocks start at line %d and end at line %d\n", inputsBlockStarts, inputsBlockEnds)
	log.Printf("But in the actual text, it corresponds to Line %d and Line %d\n", inputsBlockStarts+1, inputsBlockEnds-1)

	// Use bufio.Scanner to scan through the content line by line between the inputs block start and end
	scanner := bufio.NewScanner(strings.NewReader(strings.Join(lines[inputsBlockStarts:inputsBlockEnds-1], "\n")))
	for scanner.Scan() {
		line := scanner.Text()
		// Skip empty lines and comments
		if len(line) == 0 || strings.Contains(line, "#") {
			continue
		}

		// Split the line into key and value
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			// Check for nested objects
			if strings.Contains(value, "{") {
				// Push the current key onto the nesting stack
				nestingStack = append(nestingStack, key)

				// Initialize the nested object content buffer
				nestedObjectContent := value + "\n"

				// Keep reading lines until the closing curly brace of the nested object is found
				// This scanner only starts scanning from the next line of the line containing the opening curly brace
				for scanner.Scan() {
					nestedLine := scanner.Text()

					nestedObjectContent += nestedLine + "\n"

					// Check if the line contains the closing curly brace
					if strings.Contains(nestedLine, "}") {
						// Pop the current key from the nesting stack
						key = strings.Join(nestingStack, ".")
						// Add the nested object to the key-value pairs
						keyValuePairs = append(keyValuePairs, KeyValue{Key: key, Value: nestedObjectContent})
						break
					}
				}
			} else {
				// If there is a nesting stack, combine the key with the current nesting level
				if len(nestingStack) > 0 {
					key = strings.Join(nestingStack, ".") + "." + key
				}
				keyValuePairs = append(keyValuePairs, KeyValue{Key: key, Value: value})
			}
		}
	}

	// Print the extracted key-value pairs
	for _, kv := range keyValuePairs {
		fmt.Printf("Key: %s, Value: %s\n", kv.Key, kv.Value)
	}
}
