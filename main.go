package main

import (
	"log"

	"github.com/algo7/terragrunt-docs/pkg/parser"
)

func main() {
	res := parser.ExtractInputsFromTerragrunt("test.hcl")
	log.Println(res)
}
