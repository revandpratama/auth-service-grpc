package adapter

import (
	"fmt"

	"github.com/revandpratama/reflect/auth-service/config"
	"github.com/twmb/franz-go/pkg/kgo"
)

type franzKafka struct {
	adapter *Adapter
}

func FranzKafka() Option {
	return &franzKafka{}
}

func (k *franzKafka) Start(a *Adapter) error {

	client, err := kgo.NewClient(
		kgo.SeedBrokers(fmt.Sprintf("%v:%v", config.ENV.KafkaHost, config.ENV.KafkaPort)), // Replace with your Kafka broker address
		kgo.ClientID(config.ENV.KafkaClientID),
		kgo.DefaultProduceTopic(config.ENV.KafkaTopic), // Default topic
	)
	if err != nil {
		return err
	}

	a.Kafka = client
	k.adapter = a

	return nil
}

func (k *franzKafka) Stop() error {

	k.adapter.Kafka.Close()

	return nil
}
