package main

import (
  "html/template"
	"path/filepath"
	"net/url"

	"worshipbinder/pkg/forms"
	"worshipbinder/pkg/models"
)

type templateData struct {
  Form *forms.Form
  FormData url.Values
	FormErrors map[string]string
	Song *models.Song
	Songs      []*models.Song
  Pages []struct{
    Number int
    Active bool
  }
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := make(map[string]*template.Template)

	pages, err := filepath.Glob(filepath.Join(dir, "*.html"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "layout/*.html"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "partials/*.html"))
		if err != nil {
		}

		cache[name] = ts
	}

	return cache, nil
}
