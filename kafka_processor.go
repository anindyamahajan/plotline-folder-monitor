package main

import (
	"github.com/IBM/sarama"
	"time"
)

func newKafkaProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}
	return producer, nil
}

func sendMessageToKafka(producer sarama.SyncProducer, topic, message string) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	retryCount := 0
	backoff := initialBackoff

	for {
		if retryCount >= maxRetry {
			logger.Errorf("Failed to send message to Kafka after %d retries", maxRetry)
			return
		}

		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			logger.Errorf("Failed to send message to Kafka (retry %d): %v", retryCount+1, err)
			time.Sleep(backoff)
			retryCount++
			backoff *= backoffFactor
		} else {
			logger.Infof("Message sent to partition %d at offset %d", partition, offset)
			return
		}
	}
}
