package handlers

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
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
	b, err := json.Marshal(dataList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if _, err = io.WriteString(w, string(b)); err != nil {
		log.WithField("err", err).Error("write response error")
	}
	w.WriteHeader(http.StatusOK)
}
