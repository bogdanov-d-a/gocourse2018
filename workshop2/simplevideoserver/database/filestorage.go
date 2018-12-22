package database

import (
	"os"
	"path/filepath"
)

const contentDir = "content/"
const videoFileName = "index.mp4"

func CreateVideoFile(id string) (*os.File, error) {
	dirPath := contentDir + id
	if err := os.Mkdir(dirPath, os.ModeDir); err != nil {
		return nil, err
	}
	filePath := filepath.Join(dirPath, videoFileName)
	return os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
}
