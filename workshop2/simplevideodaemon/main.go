package main

import (
	"github.com/bogdanov-d-a/gocourse2018/workshop2/simplevideoserver/database"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

const workerCount = 3

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	if file, err := os.OpenFile("my.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666); err == nil {
		log.SetOutput(file)
	}

	db := database.Open()
	defer db.Close()

	killSignalChan := getKillSignalChan()

	var wg sync.WaitGroup
	stopPumpChan := make(chan struct{})
	stopWorkerChan := make(chan struct{}, workerCount)
	workerTaskChan := make(chan database.PendingVideoData, workerCount)

	go func() {
		wg.Add(1)
		defer wg.Done()

		for {
			select {
			case <-stopPumpChan:
				return
			default:
			}

			list, err := db.GetPendingVideoList()

			if err == nil {
				db.MarkPendingVideosAsInProgress()

				for i := 0; i < len(list); i++ {
					workerTaskChan <- list[i]
				}
			}

			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < workerCount; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()

			for {
				select {
				case <-stopWorkerChan:
					return
				case task := <-workerTaskChan:
					processVideoTask(task, db)
				}
			}
		}()
	}

	waitForKillSignal(killSignalChan)

	stopPumpChan <- struct{}{}
	for i := 0; i < workerCount; i++ {
		stopWorkerChan <- struct{}{}
	}
	wg.Wait()

	db.MarkInProgressVideosAsPending()
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

func getVideoDuration(videoPath string) (float64, error) {
	result, err := exec.Command(`ffprobe`, `-v`, `error`, `-show_entries`, `format=duration`, `-of`, `default=noprint_wrappers=1:nokey=1`, videoPath).Output()
	if err != nil {
		return 0.0, err
	}

	return strconv.ParseFloat(strings.Trim(string(result), "\n\r"), 64)
}

func ffmpegTimeFromSeconds(seconds int64) string {
	return time.Unix(seconds, 0).UTC().Format(`15:04:05.000000`)
}

func createVideoThumbnail(videoPath string, thumbnailPath string, thumbnailOffset int64) error {
	return exec.Command(`ffmpeg`, `-i`, videoPath, `-ss`, ffmpegTimeFromSeconds(thumbnailOffset), `-vframes`, `1`, thumbnailPath).Run()
}

func processVideoTask(task database.PendingVideoData, db database.Database) {
	duration := 0
	defer func() {
		db.MarkVideoAsComplete(task.DbID, duration)
	}()

	videoPath := database.DaemonPathWrap(database.GetVideoFilePath(task.ID))

	if durationFloat, err := getVideoDuration(videoPath); err == nil {
		duration = int(durationFloat)
	} else {
		return
	}

	thumbnailPath := database.DaemonPathWrap(database.GetPreviewFilePath(task.ID))
	os.Remove(thumbnailPath)
	createVideoThumbnail(videoPath, thumbnailPath, int64(duration/2))
}
