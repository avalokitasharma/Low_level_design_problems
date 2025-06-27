package loggerservice

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

type Logger struct {
	mu          sync.Mutex
	file        *os.File
	logLevel    LogLevel
	logFilePath string
	maxFileSize int64 // in bytes
	sinks       []func(string)
}

func NewLogger(filePath string, level LogLevel, maxFileSize int64) (*Logger, error) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %v", err)
	}
	return &Logger{
		file:        file,
		logLevel:    level,
		logFilePath: filePath,
		maxFileSize: maxFileSize,
		sinks:       []func(string){},
	}, nil
}

func (l *Logger) log(level LogLevel, format string, args ...interface{}) {
	if level < l.logLevel {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	message := fmt.Sprintf(format, args...)
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf("[%s] [%s] %s\n", timestamp, logLevelStrings[level], message)

	if _, err := l.file.WriteString(logEntry); err != nil {
		log.Printf("failed to write to log file: %v", err)
	}
	if l.needsRotation() {
		if err := l.rotateLogFile(); err != nil {
			log.Printf("failed to rotate log file: %v", err)
		}
	}

}

func (l *Logger) needsRotation() bool {
	info, err := l.file.Stat()
	if err != nil {
		log.Printf("failed to get log file info: %v", err)
		return false
	}
	return info.Size() >= l.maxFileSize
}

func (l *Logger) rotateLogFile() error {

	if err := l.file.Close(); err != nil {
		return fmt.Errorf("failed to close current log file: %v", err)
	}
	backupPath := fmt.Sprintf("%s.%d", l.logFilePath, time.Now().Unix())
	if err := os.Rename(l.logFilePath, backupPath); err != nil {
		return fmt.Errorf("failed to rename log file: %v", err)
	}
	file, err := os.OpenFile(l.logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open new log file: %v", err)
	}
	l.file = file
	return nil

}

func (l *Logger) ReadLogs() ([]string, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if err := l.file.Close(); err != nil {
		return nil, fmt.Errorf("failed to close log file for reading: %v", err)
	}

	file, err := os.Open(l.logFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file for reading: %v", err)
	}
	defer file.Close()

	var logs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		logs = append(logs, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading log file: %v", err)
	}
	l.file, err = os.OpenFile(l.logFilePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to reopen log file: %v", err)
	}

	return logs, nil

}

func (l *Logger) SetLogLevel(level LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logLevel = level
}

func (l *Logger) Info(format string, args ...interface{}) {
	l.log(INFO, format, args...)
}

func (l *Logger) Debug(format string, args ...interface{}) {
	l.log(DEBUG, format, args...)
}

func (l *Logger) Warning(format string, args ...interface{}) {
	l.log(WARNING, format, args...)
}

func (l *Logger) Error(format string, args ...interface{}) {
	l.log(ERROR, format, args...)
}

func (l *Logger) AddOutputSink(sink func(string)) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.sinks = append(l.sinks, sink)
}

func (l *Logger) Close() {
	l.mu.Lock()
	defer l.mu.Unlock()
	if err := l.file.Close(); err != nil {
		log.Printf("failed to close log file: %v", err)
	}
}
