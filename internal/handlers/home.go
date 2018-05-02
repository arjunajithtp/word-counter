package handlers

import (
	"net/http"
	"log"
	"github.com/arjunajithtp/word-counter/internal/helpers"
	"github.com/arjunajithtp/word-counter/internal/services"
)

func HomeHandler(w http.ResponseWriter, r *http.Request)  {

	if r.Method == http.MethodPost {
		r.ParseForm()
		url := r.Form.Get("url")
		dataByte, err := helpers.GetData(url)
		if err != nil {
			log.Println("error while trying to get webite contents: ", err)
		}
		words, err := services.HTMLTagRemover(dataByte)
		if err != nil {
			log.Println("error while trying to remove the html tags: ", err)
		}
		log.Println("Words: ", words)

		return
	}
	err := helpers.RenderPage(w, "home", nil)
	if err != nil {
		log.Println("error while trying to render page on home handler: ", err)
	}
}
