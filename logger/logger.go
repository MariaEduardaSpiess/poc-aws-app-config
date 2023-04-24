package logger

import (
	"log"
	"os"
)

func Log(msg string) {
	// Open a log file for writing
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a logger that writes to the file
	logger := log.New(file, "", log.Ldate|log.Ltime)

	logger.Println(msg)
}
