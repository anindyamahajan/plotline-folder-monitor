package main

import (
	"github.com/fsnotify/fsnotify"
)

func main() {
	initLogger()
	logger.Info("Starting the folder monitor...")

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logger.Fatal(err)
	}
	defer watcher.Close()

	producer, err := newKafkaProducer([]string{kafkaBroker})
	if err != nil {
		logger.Fatal(err)
	}
	defer producer.Close()

	done := make(chan bool)
	go watchFiles(watcher, producer)

	err = watcher.Add(watchDir)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infof("Watching directory: %s", watchDir)
	<-done
}
