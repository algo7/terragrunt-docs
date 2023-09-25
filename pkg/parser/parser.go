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
func ExtractInputsFromTerragrunt(file string) (string, error) {

	// Read the file
	content, err := os.ReadFile(file)
	utils.ErrorHandler(err)

	// Match
	inputsBlockFound := inputsBlockPattern.FindString(string(content))

	if inputsBlockFound == "" {
		return "Default Settings", nil
	}

	return inputsBlockFound, nil
}
