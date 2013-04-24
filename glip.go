package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var languages = []string{ "JavaScript", "Ruby", "Python", "Java", "Shell", "PHP", "C", "C++", "Perl", "Objective-C", "Go", "D", "Awk", "Dart", "Clojure", "Scala", "Haskell", "Lua", "Rust", "Common Lisp", "Erlang", "CoffeeScript" }

func main() {
	for _, language := range languages {
		fmt.Printf("%d %s\n", getSumForLanguage(language), language)
	}
}

func getSumForLanguage(language string) uint64 {
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

		if strings.Contains(line, "mini-icon-star") {
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
