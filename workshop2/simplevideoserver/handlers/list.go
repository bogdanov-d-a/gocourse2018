package handlers

import (
	"encoding/json"
	"github.com/bogdanov-d-a/gocourse2018/workshop2/simplevideoserver/database"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type listItemData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Duration  int    `json:"duration"`
	Thumbnail string `json:"thumbnail"`
}

func list(ids []database.VideoData, w http.ResponseWriter) {
	responseData := make([]listItemData, len(ids))
	for i := 0; i < len(ids); i++ {
		dataSrc := ids[i]
		responseData[i] = listItemData{
			dataSrc.ID,
			dataSrc.Name,
			dataSrc.Duration,
			"/" + database.GetPreviewFilePath(dataSrc.ID)}
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
