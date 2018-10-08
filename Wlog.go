package log

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	TRACK *log.Logger
	INFO *log.Logger
	WARN *log.Logger
	ERROR *log.Logger
)

func init() {
	file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	TRACK = log.New(ioutil.Discard, "TRACK: ", log.Ldate|log.Ltime|log.Lshortfile)
	INFO = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WARN = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	ERROR = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

