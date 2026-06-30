package logging

import (
	"context"
	"log/slog"
	"os"
)

func NewLogger(ctx context.Context, loggingOptions *LoggingOptions) (*slog.Logger, error) {

	file, err := os.OpenFile(
		loggingOptions.PathFile,
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

	return logger, nil
}
