package models

// import (
// 	"database/sql"
// )

type SongModel struct {
  // DB *sql.DB
}

func (s *SongModel) Insert(
  title, artirst, copyright string,
  tempo int,
  meter, sheet string,
) (int, error) {
  return 1, nil
}

func (m *SongModel) Get(id int) (*Song, error) {
  return &Song{
    ID: 1,
    Title: "Living Hope",
    Artist: "Phil Wickham",
    Copyright: "Public Domain",
    Tempo: 123,
    Meter: "1/4",
    Sheet: "Liiiiiiiviiiiinggg hoooopeeeee",
  }, nil
}

func (m *SongModel) List() ([]*Song, error) {
  songs := []*Song{
    &Song{
      ID: 1,
      Title: "Living Hope",
      Artist: "Phil Wickham",
      Copyright: "Public Domain",
      Tempo: 123,
      Meter: "1/4",
      Sheet: "Liiiiiiiviiiiinggg hoooopeeeee",
    },
    &Song{
      ID: 2,
      Title: "Waymaker",
      Artist: "Leeland",
      Copyright: "Public Domain",
      Tempo: 88,
      Meter: "1/8",
      Sheet: "Waaaaaaaymaakerr",
    },
    &Song{
      ID: 3,
      Title: "All Who Are Thirsty",
      Artist: "Brenton Brown / Glenn Robertson",
      Copyright: "Public Domain",
      Tempo: 68,
      Meter: "1/4",
      Sheet: "THIRSTYYYY",
    },
  }

  return songs, nil
}

func (m *SongModel) Count() (int, error) {
  return 100, nil
}
