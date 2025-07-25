package logger

import (
	"io"
	"log"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Info     *log.Logger
	Error    *log.Logger
	HTTP     *log.Logger
	Activity *log.Logger
)

func Init() {
	logOutput := io.MultiWriter(
		os.Stdout,
		&lumberjack.Logger{
			Filename:   "logs/app.log",
			MaxSize:    10, // MB
			MaxBackups: 3,
			MaxAge:     28,   // days
			Compress:   true, // gzip
		},
	)

	Info = log.New(logOutput, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(logOutput, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	HTTP = log.New(logOutput, "HTTP: ", log.Ldate|log.Ltime)
	Activity = log.New(logOutput, "ACTIVITY: ", log.Ldate|log.Ltime)
}
