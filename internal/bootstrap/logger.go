package bootstrap

import (
	"fmt"
	"strings"

	"go.uber.org/zap"

	"github.com/lvlBA/online_shop/pkg/logger"
)

type ConfigLogger struct {
	Level             string
	DisableCaller     bool
	DisableStacktrace bool
}

func InitLogger(cfg *ConfigLogger) (logger.Logger, error) {
	zapLoggerCfg := zap.NewProductionConfig()

	switch strings.ToLower(cfg.Level) {
	case zap.InfoLevel.String():
		zapLoggerCfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case zap.WarnLevel.String():
		zapLoggerCfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case zap.ErrorLevel.String():
		zapLoggerCfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case zap.FatalLevel.String():
		zapLoggerCfg.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	case zap.PanicLevel.String():
		zapLoggerCfg.Level = zap.NewAtomicLevelAt(zap.PanicLevel)
	default:
		zapLoggerCfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}
	zapLoggerCfg.DisableCaller = cfg.DisableCaller
	zapLoggerCfg.DisableStacktrace = cfg.DisableStacktrace

	zapLogger, err := zapLoggerCfg.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to create logger: %w", err)
	}

	return logger.NewZapLogger(zapLogger.Sugar()), nil
}
