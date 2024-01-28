package main

import (
	"github.com/IBM/sarama"
	"github.com/fsnotify/fsnotify"
	"path/filepath"
)

func watchFiles(watcher *fsnotify.Watcher, producer sarama.SyncProducer) {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Create == fsnotify.Create {
				logger.Infof("New file detected: %s", event.Name)
				if filepath.Ext(event.Name) == ".csv" {
					if validateCSVColumns(event.Name) {
						fileName := filepath.Base(event.Name)
						go sendMessageToKafka(producer, kafkaTopic, fileName)
					} else {
						logger.Errorf("ALERT::CSV validation failed for file: %s", event.Name)
					}
				} else {
					logger.Infof("Ignored non-csv file: %s", event.Name)
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			logger.Errorf("Error detected: %v", err)
		}
	}
}
