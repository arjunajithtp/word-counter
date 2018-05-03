package services

import (
	"golang.org/x/net/html"
	"strings"
	"fmt"
)

func HTMLTagRemover(dataBytes []byte) ([]string, error) {
	var words []string
	domDocTest := html.NewTokenizer(strings.NewReader(string(dataBytes)))
	previousStartTokenTest := domDocTest.Token()
	for {
		tt := domDocTest.Next()
		switch {
		case tt == html.ErrorToken:
			return words, fmt.Errorf("error notation %v is encounded", html.ErrorToken)
		case tt == html.StartTagToken:
			previousStartTokenTest = domDocTest.Token()
		case tt == html.TextToken:
			if previousStartTokenTest.Data == "script" {
				continue
			}
			word := strings.TrimSpace(html.UnescapeString(string(domDocTest.Text())))
			lastChar  := word[len(word)-1:]
			if lastChar == "." || lastChar == "," || lastChar == ";" || lastChar == "!" || lastChar == "?" || lastChar == ":" {
				word = word[:len(word)-1]
			}
			if len(word) > 0 {
				words = append(words, word)
			}
		}
	}
	return words, nil
}