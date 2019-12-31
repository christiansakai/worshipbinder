package main

import (
  "html/template"
  "net/http"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      tmpl := template.Must(template.ParseFiles(
        "./web/templates/home.page.html",
        "./web/templates/layout/base.html",
        "./web/templates/partials/header.html",
        "./web/templates/partials/footer.html",
      ))

      tmpl.Execute(w, nil)
  })

  http.HandleFunc("/songs", func(w http.ResponseWriter, r *http.Request) {
      tmpl := template.Must(template.ParseFiles(
        "./web/templates/song.list.page.html",
        "./web/templates/layout/base.html",
        "./web/templates/partials/header.html",
        "./web/templates/partials/footer.html",
        "./web/templates/partials/pagination.html",
      ))

      tmpl.Execute(w, nil)
  })

  http.HandleFunc("/songs/create", func(w http.ResponseWriter, r *http.Request) {
      tmpl := template.Must(template.ParseFiles(
        "./web/templates/song.create.page.html",
        "./web/templates/layout/base.html",
        "./web/templates/partials/header.html",
        "./web/templates/partials/footer.html",
      ))

      tmpl.Execute(w, nil)
  })

  http.HandleFunc("/songs/1", func(w http.ResponseWriter, r *http.Request) {
      tmpl := template.Must(template.ParseFiles(
        "./web/templates/song.show.page.html",
        "./web/templates/layout/base.html",
        "./web/templates/partials/header.html",
        "./web/templates/partials/footer.html",
      ))

      tmpl.Execute(w, nil)
  })

  http.HandleFunc("/songs/1/edit", func(w http.ResponseWriter, r *http.Request) {
      tmpl := template.Must(template.ParseFiles(
        "./web/templates/song.edit.page.html",
        "./web/templates/layout/base.html",
        "./web/templates/partials/header.html",
        "./web/templates/partials/footer.html",
      ))

      tmpl.Execute(w, nil)
  })

  fs := http.FileServer(http.Dir("web/static"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))

  http.ListenAndServe(":80", nil)
}
