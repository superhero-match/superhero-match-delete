package service

import (
	"github.com/superhero-match-delete/internal/config"
	"github.com/superhero-match-delete/internal/producer"
	"go.uber.org/zap"
)

// Service holds all the different services that are used when handling request.
type Service struct {
	Producer   *producer.Producer
	Logger     *zap.Logger
	TimeFormat string
}

// NewService creates value of type Service.
func NewService(cfg *config.Config) (*Service, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	defer logger.Sync()

	return &Service{
		Producer:   producer.NewProducer(cfg),
		Logger:     logger,
		TimeFormat: cfg.App.TimeFormat,
	}, nil
}
