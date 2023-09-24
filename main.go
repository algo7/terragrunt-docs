package main

import (
	"log"
	"os"
	"regexp"
)

// extractInputsFromTerragrunt extracts the inputs block from a terragrunt file
func extractInputsFromTerragrunt(file string) (string, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	// Regular expression to match the inputs block
	inputsPattern := regexp.MustCompile(`(?s)inputs\s*=\s*{.*?}`)
	match := inputsPattern.FindString(string(content))

	if match == "" {
		return "Default Settings", nil
	}

	return match, nil
}

func main() {
	res, err := extractInputsFromTerragrunt("test.hcl")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}
