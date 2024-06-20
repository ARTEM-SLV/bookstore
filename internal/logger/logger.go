package logger

import (
	"log"
	"os"
)

func writeLog(msg string, state string) {
	file, err := os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Не удалось открыть файл для логов:", err)
		return
	}
	defer file.Close()

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix(state)
	log.SetOutput(file)

	log.Println(msg)
}

func Info(msg string) {
	writeLog(msg, "INFO: ")
}

func Error(msg string) {
	writeLog(msg, "ERROR: ")
}
