package conntracker

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log/slog"
	"time"
)

func InitConfig() (Config, error) {

	cfg, err := ini.Load("/etc/conntracker/conntracker.ini")
	if err != nil {
		return Config{}, fmt.Errorf("failed to load ini file: %w", err)
	}

	// Get the default section
	defaultSection := cfg.Section("default")

	// Load TimeLocation
	timeLocationStr := defaultSection.Key("TimeLocation").String()
	loc, err := time.LoadLocation(timeLocationStr)
	if err != nil {
		return Config{}, fmt.Errorf("failed to load location: %w", err)
	}

	// Parse LogLevel
	logLevelStr := defaultSection.Key("LogLevel").String()
	logLevel, err := parseLogLevel(logLevelStr)
	if err != nil {
		return Config{}, fmt.Errorf("failed to parse log level: %w", err)
	}

	// Parse MaxWorkers with a default value of 4
	maxWorkers := defaultSection.Key("MaxWorkers").MustInt(4)

	return Config{
		AuditLogPath:    defaultSection.Key("AuditLogPath").String(),
		SqlitePath:      defaultSection.Key("SqlitePath").String(),
		TimeFormatAudit: defaultSection.Key("TimeFormatAudit").String(),
		TimeFormatDB:    defaultSection.Key("TimeFormatDB").String(),
		TimeLocation:    loc,
		MaxWorkers:      maxWorkers,
		LogLevel:        logLevel,
	}, nil
}

func parseLogLevel(level string) (slog.Level, error) {
	switch level {
	case "DEBUG":
		return slog.LevelDebug, nil
	case "INFO":
		return slog.LevelInfo, nil
	case "WARN":
		return slog.LevelWarn, nil
	case "ERROR":
		return slog.LevelError, nil
	default:
		return slog.LevelInfo, fmt.Errorf("unknown log level: %s", level)
	}
}
