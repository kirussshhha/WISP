package rabbitmq

import (
	"WISP/internal/core/domain"
	"encoding/json"
	"time"

	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

type Consumer struct {
	client *Client
}

func NewConsumer(client *Client) *Consumer {
	return &Consumer{client: client}
}

func (c *Consumer) StartConsuming() error {
	msgs, err := c.client.channel.Consume(
		c.client.config.Queue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go c.processBatches(msgs)
	return nil
}

func (c *Consumer) processBatches(msgs <-chan amqp091.Delivery) {
	var batch []domain.EmailMessage
	ticker := time.NewTicker(time.Minute)

	for {
		select {
		case msg := <-msgs:
			var email domain.EmailMessage
			if err := json.Unmarshal(msg.Body, &email); err != nil {
				log.Error().Err(err).Msg("Failed to unmarshal message")
				msg.Nack(false, true)
				continue
			}

			batch = append(batch, email)
			if len(batch) >= 10 {
				c.sendEmailBatch(batch)
				batch = batch[:0]
			}
			msg.Ack(false)

		case <-ticker.C:
			if len(batch) > 0 {
				c.sendEmailBatch(batch)
				batch = batch[:0]
			}
		}
	}
}

func (c *Consumer) sendEmailBatch(batch []domain.EmailMessage) {
	log.Info().
		Int("batch_size", len(batch)).
		Msg("Processing email batch")

	for _, msg := range batch {
		log.Info().
			Str("to", msg.To).
			Str("subject", msg.Subject).
			Msg("Email sent")
	}
}
