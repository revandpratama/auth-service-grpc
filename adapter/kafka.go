package adapter

import (
	"log"

	"github.com/twmb/franz-go/pkg/kgo"
)

type franzKafka struct {
	adapter *Adapter
}

func FranzKafka() Option {
	return &franzKafka{}
}

func (k *franzKafka) Start(a *Adapter) error {

	topics := []string{"generalauthservice"}

	client, err := kgo.NewClient(
		kgo.SeedBrokers("localhost:9092"), 
		kgo.ClientID("client-id"),
		kgo.ConsumeTopics(topics...), 
		// kgo.ConsumerGroup("group1"),       
		// kgo.AutoCommitMarks(),
		// kgo.ConsumeResetOffset(kgo.NewOffset().AtStart()),
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
	log.Println("Kafka (franz) stopped")

	return nil
}
