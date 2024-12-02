package main

import "net/http"

func (app *Application) NewServer() *http.ServeMux {
	hs := http.NewServeMux()

	return hs
}
