package main

import (
	"context"
	"github.com/bogdanov-d-a/gocourse2018/workshop2/simplevideoserver/database"
	"github.com/bogdanov-d-a/gocourse2018/workshop2/simplevideoserver/handlers"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	if file, err := os.OpenFile("my.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666); err == nil {
		log.SetOutput(file)
	}

	db := database.Open()
	defer db.Close()

	/*db.Init()
	db.AddVideo(database.VideoData{"d290f1ee-6c54-4b01-90e6-d701748f0851", "Black Retrospetive Woman", 15})
	db.AddVideo(database.VideoData{"sldjfl34-dfgj-523k-jk34-5jk3j45klj34", "Go Rally TEASER-HD", 41})
	db.AddVideo(database.VideoData{"hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345", "LMAO", 92})*/

	serverURL := ":8000"
	log.WithFields(log.Fields{"url": serverURL}).Info("starting the server")
	router := handlers.Router(db)

	killSignalChan := getKillSignalChan()
	srv := startServer(serverURL, router)
	waitForKillSignal(killSignalChan)
	srv.Shutdown(context.Background())
}

func startServer(serverURL string, handler http.Handler) *http.Server {
	srv := &http.Server{Addr: serverURL, Handler: handler}
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
	return srv
}

func getKillSignalChan() chan os.Signal {
	osKillSignalChan := make(chan os.Signal, 1)
	signal.Notify(osKillSignalChan, os.Interrupt, syscall.SIGTERM)
	return osKillSignalChan
}

func waitForKillSignal(killSignalChan <-chan os.Signal) {
	killSignal := <-killSignalChan
	switch killSignal {
	case os.Interrupt:
		log.Info("got SIGINT...")
	case syscall.SIGTERM:
		log.Info("got SIGTERM...")
	}
}
