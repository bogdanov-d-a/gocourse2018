package handlers

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
  r := mux.NewRouter()
  s := r.PathPrefix("/api/v1").Subrouter()
  s.HandleFunc("/list", list).Methods(http.MethodGet)
  s.HandleFunc("/video/d290f1ee-6c54-4b01-90e6-d701748f0851", video).Methods(http.MethodGet)
  return r
}

func list(w http.ResponseWriter, _ *http.Request) {
  fmt.Fprint(w,`[{
    "id": "d290f1ee-6c54-4b01-90e6-d701748f0851",
    "name": "Black Retrospetive Woman",
    "duration": 15,
    "thumbnail": "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg"
}]`)
}

func video(w http.ResponseWriter, _ *http.Request) {
  fmt.Fprint(w,`{
    "id": "d290f1ee-6c54-4b01-90e6-d701748f0851",
    "name": "Black Retrospetive Woman",
    "duration": 15,
    "thumbnail":"/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg",
    "url":"/content/d290f1ee-6c54-4b01-90e6-d701748f0851/index.mp4" 
}`)
}
