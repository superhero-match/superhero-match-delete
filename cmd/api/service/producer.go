package service

import (
	"github.com/superhero-match-delete/cmd/api/model"
	producer "github.com/superhero-match-delete/internal/producer/model"
)

// StoreMatch publishes new match on Kafka topic for it to be
// consumed by consumer and stored in DB.
func (s *Service) DeleteMatch(m model.Match) error {
	return s.Producer.DeleteMatch(producer.Match{
		ID: m.ID,
	})
}
