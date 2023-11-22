package handlers

import (
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
		if err := r.ParseForm(); err != nil {
			app.errors(w, http.StatusBadRequest)
			return
		}
		if err := app.Ascii.IsEngByLoop(); err != nil {
			app.errors(w, http.StatusBadGateway)
			return
		}

		err := app.Ascii.GetFs(false, r.FormValue("transformationOption"))
		if err != nil {
			app.errors(w, http.StatusBadRequest)
			return
		}

		if err = app.Ascii.DrawAscii(); err != nil {
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

		if err = tmpl.Execute(w, app); err != nil {
			app.errors(w, http.StatusInternalServerError)
		}
	} else {

		tmpl, err := template.ParseFiles("./pkg/web/html/index.html")
		if err != nil {

			app.errors(w, http.StatusInternalServerError)
			return
		}

		if err = tmpl.Execute(w, app); err != nil {
			app.errors(w, http.StatusInternalServerError)
			return
		}

	}
}

func (app *Aplication) download(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	if len(text) <= 1 {
		app.errors(w, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Disposition", "attachment; filename=file.txt")
	w.Header().Set("Content-Type", "text/plain")
	_, err := w.Write([]byte(text))
	if err != nil {
		app.errors(w, http.StatusInternalServerError)
		return
	}
}
