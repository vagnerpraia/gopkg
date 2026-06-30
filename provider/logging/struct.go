package logging

import (
	"context"
	"log/slog"
	"os"
)

type Logger struct {
	Logger  *slog.Logger
	Options *LoggingOptions
}

func NewLogger(ctx context.Context, options *LoggingOptions) (*Logger, error) {

	file, err := os.OpenFile(
		options.PathFile,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644,
	)
	if err != nil {
		return nil, err
	}

	handler := slog.NewJSONHandler(file, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	logger := slog.New(handler)

	return &Logger{
		Logger:  logger,
		Options: options,
	}, nil
}
