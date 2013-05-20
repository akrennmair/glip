package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type langsum struct {
	language string
	sum uint64
}

func main() {
	languages := findLanguages()

	data := map[string]uint64{}

	datachan := make(chan langsum)

	for _, language := range languages {
		go func(language string) {
			sum := getSumForLanguage(language)
			datachan <- langsum{language: language, sum: sum}
		}(language)
	}

	for i := 0; i < len(languages); i++ {
		langSum := <-datachan
		data[langSum.language] = langSum.sum
	}

	enc := json.NewEncoder(os.Stdout)
	enc.Encode(data)
}

func findLanguages() []string {
	resp, err := http.Get("https://github.com/languages")
	if err != nil {
		return []string{}
	}

	pattern := "<li><a href=\"/languages/"

	defer resp.Body.Close()

	languages := []string{}

	r := bufio.NewReader(resp.Body)
	for {
		line, err := r.ReadString('\n')

		if pos := strings.Index(line, pattern); pos != -1 {
			part := line[pos+len(pattern):]
			lang, _ := url.QueryUnescape(strings.Split(part, "\"")[0])
			languages = append(languages, lang)
		}

		if err != nil {
			break
		}
	}

	return languages
}

func getSumForLanguage(language string) uint64 {
	language = strings.Replace(language, "#", "%23", -1)
	url := fmt.Sprintf("https://github.com/languages/%s/most_watched", language)


	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error downloading %s: %v", url, err)
		return 0
	}
	defer resp.Body.Close()

	sum := uint64(0)
	r := bufio.NewReader(resp.Body)
	for {
		line, err := r.ReadString('\n')

		if strings.Contains(line, "octicon-star") {
			elems := strings.Split(line, " ")
			starsStr := strings.TrimRight(strings.Replace(elems[len(elems)-1], ",", "", -1), "\r\n\t ")
			stars, err := strconv.ParseUint(starsStr, 10, 64)
			if err == nil {
				sum += stars
			}
		}
		if err != nil {
			break
		}
	}
	return sum
}
