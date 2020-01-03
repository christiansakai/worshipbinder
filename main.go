package main

import (
  "html/template"
  "net/http"
)

type Song struct {
  Title string
  Artist string
  Copyright string
  Tempo string
  Meter string
  Sheet string
}

type TemplateData struct {
  Pages int
  Songs []Song
}

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

      songs := []Song{
        Song{
          Title: "Living Hope",
          Artist: "Phil Wickham",
          Copyright: "Public Domain",
          Tempo: "123",
          Meter: "1/4",
        },
        Song{
          Title: "Waymaker",
          Artist: "Leeland",
          Copyright: "Public Domain",
          Tempo: "88",
          Meter: "1/8",
        },
        Song{
          Title: "All Who Are Thirsty",
          Artist: "Brenton Brown / Glenn Robertson",
          Copyright: "Public Domain",
          Tempo: "68",
          Meter: "1/4",
        },
      }

      tmpl.Execute(w, TemplateData{
        Pages: 10,
        Songs: songs,
      })
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
