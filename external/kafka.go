package external

import (
	"context"
	"ecommerce-order/helpers"
	"fmt"
	"strings"

	"github.com/IBM/sarama"
	"github.com/pkg/errors"
)

func (e *External) ProduceKafkaMessage(ctx context.Context, topic string, data []byte) error {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	brokers := strings.Split(helpers.GetEnv("KAFKA_BROKERS", "localhost:9092,localhost:9093,localhost:9094"), ",")

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return errors.Wrap(err, "failed to comunicate with kafka brokers")
	}

	defer producer.Close()

	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(data),
	}

	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		return errors.Wrap(err, "failed to produce message to kafka")
	}

	helpers.Logger.Info(fmt.Sprintf("successfuly to produce on topic %s, partition %d, offset %d", topic, partition, offset))
	return nil
}
