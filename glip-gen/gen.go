package main

import (
	"encoding/json"
	"github.com/voxelbrain/goptions"
	"html/template"
	"log"
	"os"
	"sort"
)

type entry struct {
	Name string
	Score uint64
}

type sortableEntries []entry

func (s sortableEntries) Len() int {
	return len(s)
}

func (s sortableEntries) Less(i, j int) bool {
	return s[i].Score > s[j].Score // we want biggest to be first.
}

func (s sortableEntries) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
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

	sort.Sort(sortableEntries(entries))

	if err = tmpl.Execute(os.Stdout, entries); err != nil {
		log.Fatalf("tmpl.Execute failed: %v", err)
	}
}
