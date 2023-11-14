package handlers

import (
	"html/template"
	"net/http"
)

type formDatas struct {
	input string
	fs    string
}

func (app *Aplication) mainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.errors(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.errors(w, http.StatusMethodNotAllowed)
		return
	}
	for {
		r.ParseForm()
		i := r.FormValue("")
		i = i
		r.Method = http.MethodPost
		tmlp := template.Must(template.ParseFiles("./internal/web/html/index.html"))
		err := tmlp.ExecuteTemplate(w, "index", nil)
		if err != nil {
			app.errors(w, http.StatusNotFound)
			return
		}
	}
}
