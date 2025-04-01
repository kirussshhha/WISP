package rabbitmq

import (
    "WISP/internal/config"
    "go.uber.org/fx"
)

var Module = fx.Options(
    fx.Provide(
        func(cfg *config.RabbitMQConfig) (*Client, error) {
            return NewClient(cfg)
        },
        func(client *Client) *Producer {
            return NewProducer(client)
        },
        func(client *Client) *Consumer {
            return NewConsumer(client)
        },
    ),
    fx.Invoke(func(consumer *Consumer) error {
        return consumer.StartConsuming()
    }),
)