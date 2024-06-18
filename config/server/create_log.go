package server

import (
	"log"
	"log/slog"
	"os"
)

func CreateLogFile() {
	os.MkdirAll("./log", 0755)
	createLog("./log/error.log", "Error log!\n")
	createLog("./log/fatal.log", "Fatal log!\n")
	createLog("./log/warn.log", "Warn log!\n")
	createLog("./log/info.log", "Info log!\n")
}

func createLog(filePath, logMessage string) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		slog.Info(err.Error())
	}
	defer file.Close()

	data := []byte(logMessage)
	_, err = file.Write(data)
	if err != nil {
		slog.Info(err.Error())
	}

	log.Printf("File '%s' created successfully or already exists.", filePath)
}
