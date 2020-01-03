package main

import (
  "net/http"
  "time"
  "os"
  "path/filepath"
  "strings"

  "github.com/go-chi/chi"
  "github.com/go-chi/chi/middleware"
)

func (a *application) routes() http.Handler {
  r := chi.NewRouter()

  r.Use(middleware.RequestID)
  r.Use(middleware.RealIP)
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)
  r.Use(middleware.Timeout(60 * time.Second))

  r.Get("/", a.renderHome)
  r.Get("/songs", a.renderSongList)
  r.Get("/songs/create", a.renderCreateSongForm)
  r.Post("/songs/create", a.createSong)
  r.Get("/songs/{songID}", a.renderSong)
  r.Get("/songs/{songID}/edit", a.renderEditSongForm)
  r.Put("/songs/{songID}/edit", a.updateSong)
  r.Delete("/songs/{songID}/edit", a.deleteSong)

  workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "/web/static")
	fileServer(r, "/static", http.Dir(filesDir))

  return r
}

func fileServer(r chi.Router, path string, root http.FileSystem) {
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
