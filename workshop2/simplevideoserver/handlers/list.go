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
	ids := GetVideoList()
	var responseData [3]ListData
	for i := 0; i < len(ids); i++ {
		data_src := ids[i]
		responseData[i] = ListData{
			data_src.id,
			data_src.name,
			data_src.duration,
			"/content/" + data_src.id + "/screen.jpg"}
	}

	b, err := json.Marshal(responseData)
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
