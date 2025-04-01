package rabbitmq

import (
	"WISP/internal/config"
	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

type Client struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
	config  *config.RabbitMQConfig
}

func NewClient(cfg *config.RabbitMQConfig) (*Client, error) {
	log.Info().
		Str("host", cfg.Host).
		Str("port", cfg.Port).
		Msg("Connecting to RabbitMQ")

	conn, err := amqp091.Dial(cfg.GetURL())
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare(
		cfg.Queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:    conn,
		channel: ch,
		config:  cfg,
	}, nil
}
