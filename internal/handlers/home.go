package handlers

import (
	"net/http"
	"log"
	"github.com/arjunajithtp/word-counter/internal/helpers"
)

func HomeHandler(w http.ResponseWriter, r *http.Request)  {

	if r.Method == http.MethodPost {
		r.ParseForm()
		url := r.Form.Get("url")
		dataByte, err := helpers.GetData(url)
		if err != nil {
			log.Println("error while trying to get webite contents: ", err)
		}
		words := helpers.HTMLTagRemover(dataByte)

		wordCountMap := helpers.GetWordCount(words)
		log.Println(wordCountMap)
		return
	}
	err := helpers.RenderPage(w, "home", nil)
	if err != nil {
		log.Println("error while trying to render page on home handler: ", err)
	}
}
