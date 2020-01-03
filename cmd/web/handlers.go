package main

import (
  "net/http"

	"worshipbinder/pkg/models"
)

func (a *application) renderHome(w http.ResponseWriter, r *http.Request) {
	a.render(w, r, "home.html", nil)
}

func (a *application) renderSongList(w http.ResponseWriter, r *http.Request) {
  songs := []*models.Song{
    &models.Song{
      ID: 1,
      Title: "Living Hope",
      Artist: "Phil Wickham",
      Copyright: "Public Domain",
      Tempo: 123,
      Meter: "1/4",
      Sheet: "Liiiiiiiviiiiinggg hoooopeeeee",
    },
    &models.Song{
      ID: 2,
      Title: "Waymaker",
      Artist: "Leeland",
      Copyright: "Public Domain",
      Tempo: 88,
      Meter: "1/8",
      Sheet: "Waaaaaaaymaakerr",
    },
    &models.Song{
      ID: 3,
      Title: "All Who Are Thirsty",
      Artist: "Brenton Brown / Glenn Robertson",
      Copyright: "Public Domain",
      Tempo: 68,
      Meter: "1/4",
      Sheet: "THIRSTYYYY",
    },
  }

  data := templateData{
    Songs: songs,
    Pages: []struct{
      Number int
      Active bool
    }{
      {1, true},
      {2, false},
      {3, false},
      {4, false},
      {5, false},
    },
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
