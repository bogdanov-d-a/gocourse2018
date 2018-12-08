package handlers

import (
	"encoding/json"
	"io"
	"net/http"
)

type VideoData struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Duration  int    `json:"duration"`
	Thumbnail string `json:"thumbnail"`
	Url       string `json:"url"`
}

func Video(w http.ResponseWriter, _ *http.Request) {
	data := VideoData{
		"d290f1ee-6c54-4b01-90e6-d701748f0851",
		"Black Retrospetive Woman",
		15,
		"/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg",
		"/content/d290f1ee-6c54-4b01-90e6-d701748f0851/index.mp4"}

	b, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	io.WriteString(w, string(b))
	w.WriteHeader(http.StatusOK)
}
