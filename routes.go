package main

import (
	"net/http"
)

func routes() http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/home", homePage)

	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	return mux
}
