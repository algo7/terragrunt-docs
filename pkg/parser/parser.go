package parser

import (
	"os"
	"regexp"

	"github.com/algo7/terragrunt-docs/pkg/utils"
)

const (
	inputBlockRegex = `(?s)inputs\s*=\s*{.*?}`
)

var (
	inputsBlockPattern = regexp.MustCompile(inputBlockRegex)
)

// ExtractInputsFromTerragrunt extracts the inputs block from a terragrunt file
func ExtractInputsFromTerragrunt(file string) string {

	// Read the file
	content, err := os.ReadFile(file)
	utils.ErrorHandler(err)

	// Match
	inputsBlocks := inputsBlockPattern.FindString(string(content))

	if inputsBlocks == "" {
		return "Default Settings"
	}

	return inputsBlocks
}