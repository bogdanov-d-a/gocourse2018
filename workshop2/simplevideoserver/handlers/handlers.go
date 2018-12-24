package handlers

import (
	"github.com/bogdanov-d-a/gocourse2018/workshop2/simplevideoserver/database"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func logMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"method":     r.Method,
			"url":        r.URL,
			"remoteAddr": r.RemoteAddr,
			"userAgent":  r.UserAgent(),
		}).Info("got a new request")
		h.ServeHTTP(w, r)
	})
}

func Router(db database.Database) http.Handler {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()

	s.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		limit := 6
		if limit64, err := strconv.ParseInt(query.Get("limit"), 10, 0); err == nil {
			limit = int(limit64)
		}

		skip := 0
		if skip64, err := strconv.ParseInt(query.Get("skip"), 10, 0); err == nil {
			skip = int(skip64)
		}

		ids, err := db.GetVideoList(skip, limit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		List(ids, w)
	}).Methods(http.MethodGet)

	s.HandleFunc("/video/{ID}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["ID"]

		dataSrc, err := db.GetVideoDataByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		Video(dataSrc, w)
	}).Methods(http.MethodGet)

	s.HandleFunc("/video", func(w http.ResponseWriter, r *http.Request) { UploadVideo(db, w, r) }).Methods(http.MethodPost)
	return logMiddleware(r)
}
