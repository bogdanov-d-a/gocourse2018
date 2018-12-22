package handlers

import (
	"github.com/bogdanov-d-a/gocourse2018/workshop2/simplevideoserver/database"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
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
	s.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) { List(db, w, r) }).Methods(http.MethodGet)
	s.HandleFunc("/video/{ID}", func(w http.ResponseWriter, r *http.Request) { Video(db, w, r) }).Methods(http.MethodGet)
	s.HandleFunc("/video", UploadVideo).Methods(http.MethodPost)
	return logMiddleware(r)
}
