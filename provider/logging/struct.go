package logging

import (
	"context"
	"log/slog"
	"os"
)

type Logging struct {
	Logger  *slog.Logger
	Options *LoggingOptions
}

func NewLogging(ctx context.Context, options *LoggingOptions) (*Logging, error) {

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

	return &Logging{
		Logger:  logger,
		Options: options,
	}, nil
}
