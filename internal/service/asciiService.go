package service

import (
	"ascii-art/internal/handlers"
	"net/http"
)

func AsciiService() {
	app := &handlers.Aplication{
		
	}
	srv := &http.Server{
		Addr:    ":8080",
		Handler: app.Route(),
	}
	err := srv.ListenAndServe()
	if err != nil {
	}
}
