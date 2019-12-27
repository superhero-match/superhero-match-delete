package producer

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"

	"github.com/superhero-match-delete/internal/producer/model"
)

// DeleteMatch publishes match id on Kafka topic for it to be
// consumed by consumer and mark as deleted in DB.
func (p *Producer) DeleteMatch(m model.Match) error {
	var sb bytes.Buffer

	err := json.NewEncoder(&sb).Encode(m)
	if err != nil {
		return err
	}

	err = p.Producer.WriteMessages(
		context.Background(),
		kafka.Message{
			Value: sb.Bytes(),
		},
	)
	if err != nil {
		return err
	}

	return nil
}