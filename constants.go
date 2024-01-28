package main

import (
	"github.com/sirupsen/logrus"
	"time"
)

const (
	watchDir       = "/path/to/folder"
	kafkaBroker    = "localhost:9092"
	kafkaTopic     = "csv_file_address"
	maxRetry       = 10
	initialBackoff = 1000 * time.Millisecond
	backoffFactor  = 2
)

var (
	logger *logrus.Logger
)
