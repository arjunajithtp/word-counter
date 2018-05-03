package routers

import (
	"github.com/go-chi/chi"
	"github.com/arjunajithtp/word-counter/internal/handlers"
	"github.com/go-chi/chi/middleware"
	"os"
	"net/http"
	"path/filepath"
	"strings"
)

func GetRoutes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Route("/word-counter", func(r chi.Router) {
		r.Get("/home", handlers.HomeHandler)
		r.Post("/home", handlers.HomeHandler)
		workDir, _ := os.Getwd()
		FileServer(mux, "/public/", http.Dir(filepath.Join(workDir, "public/")))
	})

	/*mux.Get("/home", handlers.HomeHandler)
	mux.Post("/home", handlers.HomeHandler)*/

	return mux
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}