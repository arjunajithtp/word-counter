package main

import (
	"github.com/arjunajithtp/word-counter/internal/routers"
	"log"
	"net/http"
)

func main() {
	if err := http.ListenAndServe(":8181", routers.GetRoutes()); err != nil {
		log.Fatal("http.ListenAndServe: ", err)
	}
}
