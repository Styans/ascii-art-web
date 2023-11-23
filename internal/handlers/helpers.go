package handlers

import (
	"ascii-art/internal/asciiArt"
	"html/template"
	"net/http"
	"strconv"
)

type Aplication struct {
	Ascii asciiArt.ArtObjects
	Valid bool
}

func (app *Aplication) Route() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.mainPage)
	mux.HandleFunc("/ascii-art", app.asciiArt)
	mux.HandleFunc("/download", app.download)
	style := http.FileServer(http.Dir("./pkg/web/ui/"))
	mux.Handle("/static/", http.StripPrefix("/static", style))
	return mux
}

func (app *Aplication) errors(w http.ResponseWriter, problem int) {
	w.WriteHeader(problem)
	tmlp, err := template.ParseFiles("./pkg/web/html/error.html")
	if err != nil {
		http.Error(w, err.Error(), problem)
		return
	}
	e := "problem is " + strconv.Itoa(problem)
	tmlp.Execute(w, e)
}
