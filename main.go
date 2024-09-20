package main

import (
	"io"
	"log"
	"os"
	"encoding/json"
	"text/template"
)

func main() {

	// read command-line argument JSON file
	jsonPath := os.Args[1]
	jsonRaw, err := os.ReadFile(jsonPath)
	if err != nil {
		log.Fatalf("error reading values: %v", err)
	}

	var values interface{}
	err = json.Unmarshal(jsonRaw, &values)
	if err != nil {
		log.Fatalf("error unmarshalling values: %v", err)
	}

	// read a Go template from stdin
	tmplRaw, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("error reading template: %v", err)
	}

	// parse the template
	tmpl, err := template.New("template").Parse(string(tmplRaw))
	if err != nil {
		log.Fatalf("error parsing template: %v", err)
	}

	// execute the template and write to stdout
	err = tmpl.Execute(os.Stdout, values)
	if err != nil {
		log.Fatalf("error executing template: %v", err)
	}
}

