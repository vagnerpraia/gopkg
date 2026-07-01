package gpslog

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"
)

type Logging struct {
	logger  *slog.Logger
	options *LoggingOptions
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

	logger := slog.New(slog.NewJSONHandler(file, handlerOptions))

	return &Logging{
		logger:  logger,
		options: options,
	}, nil
}

func (that *Logging) Info(msg string) {

	if that.options.Print {
		log.Println(msg)
	}

	if that.options.WriteFile {
		that.logger.Info(msg)
	}
}

func (that *Logging) Error(msg string, err error) {

	formatted := fmt.Sprintf("%s: %v\n", msg, err)

	if that.options.Print {
		log.Println(formatted)
	}

	if that.options.WriteFile {
		that.logger.Error(formatted)
	}
}
