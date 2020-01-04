package main

import (
  "log"
  "html/template"
  "os"
  "net/http"

  // _ "github.com/lib/pq"

	"worshipbinder/pkg/models"
)

type application struct {
  debug bool
	errorLog      *log.Logger
	infoLog       *log.Logger
  templateCache map[string]*template.Template
	songs      interface {
		Insert(string, string, string, int, string, string) (int, error)
		Get(int) (*models.Song, error)
		List() ([]*models.Song, error)
		Count() (int, error)
	}
}

func main() {
  debug := true
  addr := ":4000"
  // dsn := "postgres://jesusisking:isaiah96@localhost/worshipbinder"

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// db, err := openDB(*dsn)
	// if err != nil {
		// errorLog.Fatal(err)
	// }

	// defer db.Close()

	templateCache, err := newTemplateCache("./web/templates/")
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
    debug: debug,
		errorLog:      errorLog,
		infoLog:       infoLog,
		templateCache: templateCache,
		songs:         &models.SongModel{},
	}

	infoLog.Printf("Starting server on %s", addr)

  r := app.routes()

  err = http.ListenAndServe(addr, r)
  if err != nil {
    errorLog.Fatal(err)
  }
}

// func openDB(dsn string) (*sql.DB, error) {
// 	db, err := sql.Open("postgresql", dsn)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err = db.Ping(); err != nil {
// 		return nil, err
// 	}

// 	return db, nil
// }
