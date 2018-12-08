package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
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

func Video(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]
	data_src := GetVideoListDataById(id)

	data := VideoData{
		id,
		data_src.name,
		data_src.duration,
		"/content/" + id + "/screen.jpg",
		"/content/" + id + "/index.mp4"}

	b, err := json.Marshal(data)
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
