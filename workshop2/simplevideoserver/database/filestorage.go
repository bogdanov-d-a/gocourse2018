package database

import (
	"io"
	"os"
	"path/filepath"
)

const contentDir = "content/"
const videoFileName = "index.mp4"
const previewFileName = "screen.jpg"
const defaultPreviewFilePath = "screen-default.jpg"

func getContentDirForId(id string) string {
	return contentDir + id
}

func GetVideoFilePath(id string) string {
	return filepath.Join(getContentDirForId(id), videoFileName)
}

func GetPreviewFilePath(id string) string {
	return filepath.Join(getContentDirForId(id), previewFileName)
}

func MakeContentDir(id string) error {
	return os.Mkdir(getContentDirForId(id), os.ModeDir)
}

func openNewFile(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_CREATE|os.O_WRONLY, os.ModePerm)
}

func CreateVideoFile(id string) (*os.File, error) {
	return openNewFile(GetVideoFilePath(id))
}

func CreatePreviewFileFromDefault(id string) error {
	srcFile, err := os.Open(defaultPreviewFilePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := openNewFile(GetPreviewFilePath(id))
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}
