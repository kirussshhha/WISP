package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type RabbitMQConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Queue    string
	VHost    string
}

var Module = fx.Options(
	fx.Provide(
		NewDBConfig,
		NewRabbitMQConfig,
	),
)

func NewDBConfig() *DBConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
}

func NewRabbitMQConfig() *RabbitMQConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &RabbitMQConfig{
		Host:     os.Getenv("RABBITMQ_HOST"),
		Port:     os.Getenv("RABBITMQ_PORT"),
		User:     os.Getenv("RABBITMQ_USER"),
		Password: os.Getenv("RABBITMQ_PASSWORD"),
		Queue:    os.Getenv("RABBITMQ_QUEUE"),
		VHost:    os.Getenv("RABBITMQ_VHOST"),
	}
}

func (c *RabbitMQConfig) GetURL() string {
	return "amqp://" + c.User + ":" + c.Password + "@" + c.Host + ":" + c.Port + c.VHost
}
