package server

import "net/http"

var Router = http.NewServeMux()

func StartServer() {
	server := http.Server{
		Addr:    ":8080",
		Handler: Router,
	}
	server.ListenAndServe()
}
