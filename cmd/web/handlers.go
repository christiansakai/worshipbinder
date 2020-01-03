package main

import (
  "net/http"
)

func (a *application) renderHome(w http.ResponseWriter, r *http.Request) {
	a.render(w, r, "home.html", &templateData{})
}

func (a *application) renderSongList(w http.ResponseWriter, r *http.Request) {
	a.render(w, r, "song.list.html", &templateData{})
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
