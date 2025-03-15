package logger

import (
    "log"
    "os"
)

// Global logger instance
var (
    logFile  *os.File
    logInfo  *log.Logger
    logError *log.Logger
)

// Initialize logger
func InitLogger() {
    var err error
    logFile, err = os.OpenFile("internal/lib/logger/log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("Failed to open log file: %v", err)
    }

    logInfo = log.New(logFile, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
    logError = log.New(logFile, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)

    log.Println("Logger initialized successfully")
}

// LogInfo writes informational messages
func LogInfo(format string, args ...interface{}) {
    logInfo.Printf(format, args...)
}

// LogError writes error messages
func LogError(format string, args ...interface{}) {
    logError.Printf(format, args...)
}

// CloseLogger closes the log file
func CloseLogger() {
    if logFile != nil {
        logFile.Close()
    }
}
