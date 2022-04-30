package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {

	r := mux.NewRouter()

	// define an endpoint
	r.HandleFunc("/process_text", TextHandler).Methods(http.MethodPost)

	return r
}
