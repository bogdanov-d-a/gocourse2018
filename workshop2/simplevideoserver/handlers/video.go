package handlers

import (
	"encoding/json"
	"github.com/bogdanov-d-a/gocourse2018/workshop2/simplevideoserver/database"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type videoData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Duration  int    `json:"duration"`
	Thumbnail string `json:"thumbnail"`
	URL       string `json:"url"`
}

func video(dataSrc database.VideoData, w http.ResponseWriter) {
	data := videoData{
		dataSrc.ID,
		dataSrc.Name,
		dataSrc.Duration,
		"/" + database.GetPreviewFilePath(dataSrc.ID),
		"/" + database.GetVideoFilePath(dataSrc.ID)}

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
