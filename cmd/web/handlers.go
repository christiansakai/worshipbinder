package main

import (
  "net/http"
  "fmt"
)

func (a *application) renderHome(w http.ResponseWriter, r *http.Request) {
	a.render(w, r, "home.html", nil)
}

func (a *application) renderSongList(w http.ResponseWriter, r *http.Request) {
  songs, err := a.songs.List()
	if err != nil {
		a.serverError(w, err)
		return
	}

  totalSongsCount, err := a.songs.Count() 
	if err != nil {
		a.serverError(w, err)
		return
	}

  songsPerPage := 10
  pageCount := totalSongsCount / songsPerPage
  currPage := 1

  pages := make([]struct{
    Number int
    Active bool
  }, pageCount) 

  for i := 0; i < pageCount; i++ {
    if i == currPage {
      pages[i] = struct{
        Number int
        Active bool
      }{i + 1, true}
    } else {
      pages[i] = struct{
        Number int
        Active bool
      }{i + 1, false}
    }
  }

  data := templateData{
    Songs: songs,
    Pages: pages,
  }

	a.render(w, r, "song.list.html", &data)
}

func (a *application) renderCreateSongForm(w http.ResponseWriter, r *http.Request) {
	a.render(w, r, "song.create.html", &templateData{})
}

func (a *application) createSong(w http.ResponseWriter, r *http.Request) {
	a.render(w, r, "song.create.html", &templateData{})
}

func (a *application) renderSong(w http.ResponseWriter, r *http.Request) {
	a.render(w, r, "song.show.html", &templateData{})
}

func (a *application) renderEditSongForm(w http.ResponseWriter, r *http.Request) {
	a.render(w, r, "song.edit.html", &templateData{})
}

func (a *application) updateSong(w http.ResponseWriter, r *http.Request) {
	a.render(w, r, "song.edit.html", &templateData{})
}

func (a *application) deleteSong(w http.ResponseWriter, r *http.Request) {
	a.render(w, r, "song.edit.html", &templateData{})
}
