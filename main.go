package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/algo7/terragrunt-docs/pkg/parser"
)

func main() {

	err := filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip hidden directories and any files within them
		if strings.Contains(path, "/.") {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// Parse all `terragrunt.hcl` files
		if strings.HasSuffix(info.Name(), "terragrunt.hcl") {
			inputs := parser.ExtractInputsFromTerragrunt(path)
			log.Printf("Inputs from %s:\n%s\n%s\n", path, inputs, strings.Repeat("-", 40))
		}
		return nil
	})

	if err != nil {
		log.Printf("Error walking the path %v: %v\n", "./", err)
	}

}
