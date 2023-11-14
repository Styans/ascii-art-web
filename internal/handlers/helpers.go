package handlers

import (
	"html/template"
	"net/http"
	"strconv"
)

type Aplication struct{}

func (app *Aplication) Route() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.mainPage)
	style := http.FileServer(http.Dir("./internal/web/ui/"))
	mux.Handle("/static/", http.StripPrefix("/static", style))
	return mux
}

func (app *Aplication) errors(w http.ResponseWriter, problem int) {
	w.WriteHeader(problem)
	tmlp := template.Must(template.ParseFiles("./internal/web/html/error.html"))
	e := "problem is " + strconv.Itoa(problem)
	tmlp.Execute(w, e)
}
