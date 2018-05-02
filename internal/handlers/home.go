package handlers

import (
	"net/http"
	"log"
	"github.com/arjunajithtp/word-counter/internal/helpers"
)

func HomeHandler(w http.ResponseWriter, _ *http.Request)  {
	err := helpers.RenderPage(w, "home", nil)
	if err != nil {
		log.Println("error while trying to render page on home handler: ", err)
	}
}
