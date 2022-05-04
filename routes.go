package main

import (
	"net/http"
)

func route() http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("home", homePage)

	return mux
}
