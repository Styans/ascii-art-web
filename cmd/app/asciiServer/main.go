package main

import (
	"ascii-art/internal/asciiArt"
	"ascii-art/internal/handlers"
	"fmt"
	"net/http"
)

func main() {
	app := &handlers.Aplication{
		Ascii: asciiArt.ArtObjects{},
	}
	srv := &http.Server{
		Addr:    ":8080",
		Handler: app.Route(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
