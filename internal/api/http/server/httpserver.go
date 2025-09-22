package server

import (
	"net/http"

	"github.com/ostackr/oort/internal/environment"
)

var Router = http.NewServeMux()
var Port = environment.Get("OORT_HTTP_SERVER_PORT")

func StartServer() {
	server := http.Server{
		Addr:    ":" + Port,
		Handler: Router,
	}
	server.ListenAndServe()
}
