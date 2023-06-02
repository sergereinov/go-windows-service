package logger

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	*log.Logger
	file *os.File
}

func New(path string, logToConsole bool) *Logger {
	lg := &Logger{
		Logger: log.Default(),
	}

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	lg.file = file

	if logToConsole {
		mw := io.MultiWriter(os.Stdout, lg.file)
		lg.Logger.SetOutput(mw)
	} else {
		lg.Logger.SetOutput(lg.file)
	}

	return lg
}

func (l *Logger) Close() {
	if l == nil {
		return
	}
	l.file.Close()
}
