package helpers

import (
	"golang.org/x/net/html"
	"regexp"
	"strings"
)

func HTMLTagRemover(dataBytes []byte) []string {
	var words []string
	domDocTest := html.NewTokenizer(strings.NewReader(string(dataBytes)))
	previousStartTokenTest := domDocTest.Token()
	for {
		tt := domDocTest.Next()
		switch {
		case tt == html.ErrorToken:
			return words
		case tt == html.StartTagToken:
			previousStartTokenTest = domDocTest.Token()
		case tt == html.TextToken:
			if previousStartTokenTest.Data == "script" {
				continue
			}
			token := strings.TrimSpace(html.UnescapeString(string(domDocTest.Text())))

			if len(token) > 0 {
				tokenSlice := strings.Split(token, " ")
				reg, err := regexp.Compile("[^a-zA-Z0-9]+")
				if err != nil {
					continue
				}
				for _, tokenWord := range tokenSlice {
					regSlice := reg.Split(tokenWord, -1)
					for _, word := range regSlice {
						if len(word) > 0 {
							words = append(words, word)
						}
					}
				}

			}
		}
	}
	return words
}

func GetWordCount(words []string) map[string]int {
	var wordCountMap map[string]int
	for _, word := range words {
		if wordCountMap == nil {
			wordCountMap = make(map[string]int)
		}
		count := wordCountMap[word]
		count++
		wordCountMap[word] = count
	}

	return wordCountMap
}
