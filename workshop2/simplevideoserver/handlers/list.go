package handlers

import (
	"encoding/json"
	"io"
	"net/http"
)

type ListData struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Duration  int    `json:"duration"`
	Thumbnail string `json:"thumbnail"`
}

func List(w http.ResponseWriter, _ *http.Request) {
	data := ListData{
		"d290f1ee-6c54-4b01-90e6-d701748f0851",
		"Black Retrospetive Woman",
		15,
		"/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg"}

	dataList := [1]ListData{data}
	b, _ := json.Marshal(dataList)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	io.WriteString(w, string(b))
	w.WriteHeader(http.StatusOK)
}
