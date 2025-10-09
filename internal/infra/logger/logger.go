package logger

import (
	"desafio-picpay-go2/internal/config"
	"github.com/charmbracelet/log"
	"os"
	"time"
)

func NewLogger(cfg *config.Config) *log.Logger {
	logger := log.NewWithOptions(
		os.Stdout,
		log.Options{
			TimeFormat:      time.Kitchen,
			Formatter:       log.JSONFormatter,
			ReportTimestamp: true,
		},
	)

	if cfg.Debug {
		logger.SetLevel(log.DebugLevel)
		log.SetReportCaller(true)
	}

	// In development environment, use TextFormatter for easier reading
	if cfg.Environment == "development" {
		logger.SetFormatter(log.TextFormatter)
	}
	return logger
}
