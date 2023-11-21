package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

type formDatas struct {
	input string
	fs    string
}

func (app *Aplication) mainPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	app.Ascii.Result = ""
	if r.URL.Path != "/" {
		app.errors(w, http.StatusNotFound)
		return
	}
	if r.URL.Query().Get("text") {
		text := r.FormValue("text")
		w.Header().Set("Content-Disposition", "attachment; filename=file.txt")
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(text))
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		app.Ascii.Text = r.FormValue("text")
		if err := app.Ascii.IsEngByLoop(); err != nil {
			app.errors(w, http.StatusBadRequest)
			return
		}

		app.Ascii.GetFs(false, r.FormValue("transformationOption"))
		err := app.Ascii.DrawAscii()
		if err != nil {
			app.errors(w, http.StatusBadGateway)
			return
		}

		tmpl, err := template.ParseFiles("./pkg/web/html/index.html")
		if err != nil {
			app.errors(w, http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, app.Ascii.Result)
		if err != nil {
			fmt.Println("asd")
			app.errors(w, http.StatusNotFound)
		}
	} else {

		tmpl, err := template.ParseFiles("./pkg/web/html/index.html")
		if err != nil {

			app.errors(w, http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, app.Ascii.Result)
		if err != nil {
			fmt.Println("assd")

			app.errors(w, http.StatusNotFound)
		}

	}
}
