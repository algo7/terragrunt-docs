package main

import (
	"log"
)

func main() {
	res, err := extractInputsFromTerragrunt("test.hcl")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}
