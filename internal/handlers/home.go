package handlers

import (
	"net/http"
	"log"
	"github.com/arjunajithtp/word-counter/internal/helpers"
	"encoding/json"
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
		mapByte, err := json.Marshal(wordCountMap)
		if err != nil {
			log.Println("error while trying to marshal the map data: ", err)
		}
		_, err = w.Write(mapByte)
		if err != nil {
			log.Println("error while trying to write data to UI: ", err)
		}
		return
	}
	err := helpers.RenderPage(w, "home", nil)
	if err != nil {
		log.Println("error while trying to render page on home handler: ", err)
	}
}
