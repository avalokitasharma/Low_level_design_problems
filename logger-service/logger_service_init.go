package loggerservice

import (
	"fmt"
	"log"
)

func Init() {
	logger, err := NewLogger("logger-service/app.log", INFO, 10*1024) // 10 KB max file size
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)

	}
	defer logger.Close()
	logger.Info("Application started")
	logger.Debug("Debugging information")
	logger.Warning("A warning message")
	logger.Error("An error occurred")

	logger.SetLogLevel(DEBUG)
	logger.Debug("This debug message should now be visible")

	logs, err := logger.ReadLogs()

	if err != nil {
		log.Fatalf("failed to read logs: %v", err)
	}
	fmt.Println("Logs:")
	for _, log := range logs {
		fmt.Println(log)
	}

	logger.AddOutputSink(func(entry string) {
		fmt.Println("Sink received log entry:", entry)
	})

	logger.Info("Log with custom output sink")
}
