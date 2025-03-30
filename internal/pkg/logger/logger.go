package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

func NewLogger() (*zerolog.Logger, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get working directory: %w", err)
	}

	logsDir := filepath.Join(currentDir, "logs")
	if err := os.MkdirAll(logsDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create logs directory: %w", err)
	}

	logFile := filepath.Join(logsDir, "app.log")
	file, err := os.OpenFile(
		logFile,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create/open log file: %w", err)
	}

	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
		NoColor:    false,
	}

	multi := io.MultiWriter(consoleWriter, file)

	logger := zerolog.New(multi).
		Level(zerolog.DebugLevel).
		With().
		Timestamp().
		Caller().
		Logger()

	return &logger, nil
}

func Initialize(logger *zerolog.Logger) {
	log.Logger = *logger
	logger.Info().Msg("Logger initialized")
}

var Module = fx.Options(
	fx.Provide(NewLogger),
	fx.Invoke(Initialize),
)
