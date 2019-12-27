package config

// Producer holds the configuration values for the Kafka producer.
type Producer struct {
	Brokers      []string `env:"KAFKA_BROKERS" default:"[192.168.178.17:9092]"`
	Topic        string   `env:"KAFKA_DELETE_MATCH_MATCH_TOPIC" default:"delete.match.match"`
	BatchSize    int      `env:"KAFKA_BATCH_SIZE" default:"1"`
	BatchTimeout int      `env:"KAFKA_BATCH_TIMEOUT" default:"10"`
}
