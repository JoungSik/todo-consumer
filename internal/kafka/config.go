package kafka

import "github.com/Shopify/sarama"

type KafkaConfig struct {
	Brokers []string
	Topics  []string
	Group   string
	Config  *sarama.Config
}

func (KafkaConfig) Init() *KafkaConfig {
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategyRange}
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	return &KafkaConfig{Brokers: []string{"localhost:9091"}, Group: "sarama", Topics: []string{"sarama"}, Config: config}
}
