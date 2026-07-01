package gpslog

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
)

type Logging struct {
	Logger  *slog.Logger
	Options *LoggingOptions
}

func NewLogging(ctx context.Context, options *LoggingOptions) (*Logging, error) {

	if err := os.MkdirAll(filepath.Dir(options.PathFile), 0755); err != nil {
		return nil, err
	}

	file, err := os.OpenFile(options.PathFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	handlerOptions := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	handler := slog.NewJSONHandler(file, handlerOptions)

	logger := slog.New(handler)

	return &Logging{
		Logger:  logger,
		Options: options,
	}, nil
}
