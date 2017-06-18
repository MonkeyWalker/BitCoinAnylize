package logger

import (
	// "os"
	//"io"
	"log"
	"os"
	"io"
)


func Error(e interface{})  {
	file, err := os.OpenFile("Error.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open loggerr file", ":", err)
	}

	multi := io.MultiWriter(file, os.Stdout)
	logger :=  log.New(multi,
		"",
		log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(e)
}