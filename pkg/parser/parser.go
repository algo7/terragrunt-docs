package parser

import (
	"os"
	"regexp"

	"github.com/algo7/terragrunt-docs/pkg/utils"
)

const (
	inputsBlockRegex = `(?s)inputs\s*=\s*{\s*(.*?)\s*}`
)

var (
	inputsBlockPattern = regexp.MustCompile(inputsBlockRegex)
)

// ExtractInputsFromTerragrunt extracts the content inside the inputs block from a terragrunt file
func ExtractInputsFromTerragrunt(file string) string {

	// Read the file
	content, err := os.ReadFile(file)
	utils.ErrorHandler(err)

	// Match
	matches := inputsBlockPattern.FindStringSubmatch(string(content))
	if len(matches) > 1 {
		return matches[1]
	}

	return "Default Settings"
}
