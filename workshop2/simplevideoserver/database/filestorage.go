package database

import (
	"os"
	"path/filepath"
)

const contentDir = "content/"
const videoFileName = "index.mp4"
const previewFileName = "screen.jpg"

func getContentDirForId(id string) string {
	return contentDir + id
}

func GetVideoFilePath(id string) string {
	return filepath.Join(getContentDirForId(id), videoFileName)
}

func GetPreviewFilePath(id string) string {
	return filepath.Join(getContentDirForId(id), previewFileName)
}

func makeContentDir(id string) error {
	return os.Mkdir(getContentDirForId(id), os.ModeDir)
}

func CreateVideoFile(id string) (*os.File, error) {
	if err := makeContentDir(id); err != nil {
		return nil, err
	}
	return os.OpenFile(GetVideoFilePath(id), os.O_CREATE|os.O_WRONLY, os.ModePerm)
}
