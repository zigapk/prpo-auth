package logger

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/rs/zerolog"
	"github.com/zigapk/prpo-auth/internal/config"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log zerolog.Logger

func Init() {
	var writers []io.Writer

	if config.Logs.ConsoleLogging() {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}

	if config.Logs.FileLogging() {
		writers = append(writers, newRollingFile())
	}

	mulWriter := io.MultiWriter(writers...)

	Log = zerolog.New(mulWriter).Level(zerolog.Level(config.Logs.LogLevel())).With().Timestamp().Caller().Logger()
}

func newRollingFile() io.Writer {
	dir := filepath.Dir(config.Logs.LogPath())
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	return &lumberjack.Logger{
		Filename:   config.Logs.LogPath(),
		MaxSize:    config.Logs.MaxSize(),
		MaxAge:     config.Logs.MaxAge(),
		MaxBackups: config.Logs.MaxBackups(),
	}
}
