package main

import (
	"fmt"
  "bytes"
	"net/http"
	"runtime/debug"
)

func (a *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := a.templateCache[name]
	if !ok {
		a.serverError(w, fmt.Errorf("The template %s does not exist", name))
		return
	}

	buf := new(bytes.Buffer)

	err := ts.Execute(buf, td)
	if err != nil {
		a.serverError(w, err)
		return
	}

	buf.WriteTo(w)
}

func (a *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	a.errorLog.Output(2, trace)

  if a.debug {
    http.Error(w, trace, http.StatusInternalServerError)
    return
  }

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
