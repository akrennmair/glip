package main

import (
	"encoding/json"
	"github.com/voxelbrain/goptions"
	"html/template"
	"log"
	"os"
)

type entry struct {
	Name string
	Score uint64
}

func main() {
	options := struct {
		Template string `goptions:"-t, --template, obligatory, description='Template file to render results'"`
	}{}
	goptions.ParseAndFail(&options)

	tmpl, err := template.ParseFiles(options.Template)
	if err != nil {
		log.Fatalf("Loading %s failed: %v", options.Template, err)
	}

	dec := json.NewDecoder(os.Stdin)
	data := map[string]uint64{}
	dec.Decode(&data)

	entries := []entry{}

	for language, score := range data {
		entries = append(entries, entry{Name: language, Score: score})
	}

	// TODO: sort entries

	if err = tmpl.Execute(os.Stdout, entries); err != nil {
		log.Fatalf("tmpl.Execute failed: %v", err)
	}
}
