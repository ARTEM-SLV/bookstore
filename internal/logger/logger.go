package logger

import (
	"log"
	"os"
)

type Logger struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	logFile     *os.File
}

var Log *Logger

func InitLogger() {
	var l Logger

	file, err := os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		l = Logger{
			infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
			warnLogger:  log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile),
			errorLogger: log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
			logFile:     file,
		}
	} else {
		l = Logger{
			infoLogger:  log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
			warnLogger:  log.New(file, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile),
			errorLogger: log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
			logFile:     file,
		}
	}

	Log = &l
}

// Info - метод для логирования информационных сообщений
func (l *Logger) Info(msg string) {
	go l.infoLogger.Println(msg)
}

// Warn - метод для логирования предупреждений
func (l *Logger) Warn(msg string) {
	go l.warnLogger.Println(msg)
}

// Error - метод для логирования ошибок
func (l *Logger) Error(msg string) {
	go l.errorLogger.Println(msg)
}

// CloseLog - метод для закрытия файла
func (l *Logger) CloseLog() {
	err := l.logFile.Close()
	if err != nil {
		log.Printf("File close error: %s\n", err)
	}
	log.Println("Logger file closed.")
}
