package handlers

import (
	"encoding/json"
	"github.com/bogdanov-d-a/gocourse2018/workshop2/simplevideoserver/database"
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

func List(ids []database.VideoData, w http.ResponseWriter) {
	responseData := make([]ListData, len(ids))
	for i := 0; i < len(ids); i++ {
		data_src := ids[i]
		responseData[i] = ListData{
			data_src.Id,
			data_src.Name,
			data_src.Duration,
			"/" + database.GetPreviewFilePath(data_src.Id)}
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
