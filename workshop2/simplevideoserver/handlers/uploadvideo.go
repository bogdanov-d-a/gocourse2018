package handlers

import (
	"github.com/bogdanov-d-a/gocourse2018/workshop2/simplevideoserver/database"
	"github.com/google/uuid"
	"io"
	"net/http"
)

func UploadVideo(db database.Database, w http.ResponseWriter, r *http.Request) {
	fileReader, header, err := r.FormFile("file[]")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if header.Header.Get("Content-Type") != "video/mp4" {
		http.Error(w, "Content type is not valid", http.StatusInternalServerError)
		return
	}

	id := uuid.New().String()

	if err := database.MakeContentDir(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, err := database.CreateVideoFile(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if _, err = io.Copy(file, fileReader); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := database.CreatePreviewFileFromDefault(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	db.AddVideo(database.VideoData{id, header.Filename, 0})
}
