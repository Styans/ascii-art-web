package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func (app *Aplication) mainPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	app.Ascii.Result = ""
	if r.URL.Path != "/" {
		app.errors(w, http.StatusNotFound)
		return
	}
	app.Ascii.Text = r.FormValue("text")
	if r.Method == http.MethodPost && len(app.Ascii.Text) > 0 {

		r.ParseForm()
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
		app.Valid = true

		tmpl, err := template.ParseFiles("./pkg/web/html/index.html")
		if err != nil {
			app.errors(w, http.StatusInternalServerError)
			return
		}
		// r.URL.Query()
		err = tmpl.Execute(w, app)
		if err != nil {
			app.errors(w, http.StatusNotFound)
		}
	} else {

		tmpl, err := template.ParseFiles("./pkg/web/html/index.html")
		if err != nil {

			app.errors(w, http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, app)
		if err != nil {
			fmt.Println("assd")

			app.errors(w, http.StatusNotFound)
		}

	}
}

func (app *Aplication) download(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	if len(text) <= 1 {
		app.errors(w, http.StatusInternalServerError)
		return
	}
	fmt.Println(text, len(text))
	w.Header().Set("Content-Disposition", "attachment; filename=file.txt")
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(text))
}
