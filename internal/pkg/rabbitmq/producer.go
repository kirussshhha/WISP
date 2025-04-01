package rabbitmq

import (
	"WISP/internal/core/domain"
	"encoding/json"
	"github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	client *Client
}

func NewProducer(client *Client) *Producer {
	return &Producer{client: client}
}

func (p *Producer) PublishEmails(messages []domain.EmailMessage) error {
	for _, msg := range messages {
		body, err := json.Marshal(msg)
		if err != nil {
			return err
		}

		err = p.client.channel.Publish(
			"",
			p.client.config.Queue,
			false,
			false,
			amqp091.Publishing{
				ContentType: "application/json",
				Body:        body,
			})
		if err != nil {
			return err
		}
	}
	return nil
}
