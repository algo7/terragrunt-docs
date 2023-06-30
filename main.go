package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hclparse"
)

type foo struct {
	A string   `hcl:"a, omitempty"`
	B []string `hcl:"b" omitempty"`
}

func main() {
	parser := hclparse.NewParser()

	f, parseDiags := parser.ParseHCLFile("test.hcl")

	if parseDiags.HasErrors() {
		fmt.Println(parseDiags.Error())
	}

	var fooInstance foo
	decodeDiags := gohcl.DecodeBody(f.Body, nil, &fooInstance)
	if decodeDiags.HasErrors() {
		log.Fatal(decodeDiags.Error())
	}

	fmt.Printf("%#v", fooInstance)
}
