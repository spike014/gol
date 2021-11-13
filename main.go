package gol

import (
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

const HoursOfDay = 24

var LogPath = ""

func NewLogger(types string) io.Writer {
	if LogPath == "" {
		LogPath = getLogPath()
	}

	if !pathExists(LogPath) {
		err := os.MkdirAll(LogPath, os.ModeDir|os.ModePerm)
		if err != nil {
			log.Println("MkdirAll log file path ERROR:", err)
		}
	}
	fileName := path.Join(LogPath, types+".log")
	src, err := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		os.ModeAppend|os.ModePerm,
	)
	if err != nil {
		log.Println("Open log file ERROR:", err)
	}
	defer src.Close()

	logWriter, err := rotatelogs.New(
		fileName+".%Y-%m-%d.log",
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(-1),
		rotatelogs.WithRotationTime(HoursOfDay*time.Hour),
	)
	if err != nil {
		log.Println("New rotatelogs ERROR:", err)
	}
	return logWriter
}

func SetLogger(types string, writers ...io.Writer) (writer io.Writer) {
	if types == "debug" {
		writer = io.MultiWriter(os.Stdout, io.MultiWriter(writers...))
	} else if types == "release" {
		writer = io.MultiWriter(writers...)
	}
	return
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return os.IsExist(err)
}

// getExcutePath 获取程序所在目录
func getExcutePath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		log.Println("Get program path ERROR:", err)
		return "", err
	}
	return filepath.Dir(ex), nil
}

func getLogPath() string {
	path, _ := getExcutePath()
	return filepath.Join(path, "log")
}
