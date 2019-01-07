package util

import (
	"log"
	"os"
)

var logger = log.New(os.Stderr, "  ", log.Ldate|log.Ltime)

// LogDebug adds a "DEBUG:" prefix to the log output
func LogDebug(v ...interface{}) {
	logger.SetPrefix("DEBUG:")
	logger.Println(v...)
}

// LogInfo adds a "INFO:" prefix to the log output
func LogInfo(v ...interface{}) {
	logger.SetPrefix("INFO: ")
	logger.Println(v...)
}

// LogFatal adds a "FATAL:" prefix to the log output
func LogFatal(v ...interface{}) {
	logger.SetPrefix("FATAL: ")
	logger.Println(v...)
}
