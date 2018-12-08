package handlers

import (
	"net/http"
  "github.com/gorilla/mux"
)

func Router() *mux.Router {
  r := mux.NewRouter()
  s := r.PathPrefix("/api/v1").Subrouter()
  s.HandleFunc("/list", List).Methods(http.MethodGet)
  s.HandleFunc("/video/d290f1ee-6c54-4b01-90e6-d701748f0851", Video).Methods(http.MethodGet)
  return r
}
