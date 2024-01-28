package main

import (
	"encoding/csv"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func initLogger() {
	logger = logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339Nano,
	})
}

func validateCSVColumns(fileName string) bool {
	file, err := os.Open(fileName)
	if err != nil {
		logger.Errorf("Failed to open file: %s, error: %v", fileName, err)
		return false
	}
	defer file.Close()

	reader := csv.NewReader(file)
	columns, err := reader.Read()
	if err != nil {
		logger.Errorf("Failed to read the CSV file: %s, error: %v", fileName, err)
		return false
	}

	if len(columns) != 2 || columns[0] != "userID" || columns[1] != "alertMessage" {
		logger.Errorf("CSV file %s does not have the required columns 'userID' and 'alertMessage'", fileName)
		return false
	}

	return true
}
