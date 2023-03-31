package logger

import (
	"context"

	"go.uber.org/zap"
)

type ZapLogger struct {
	log *zap.SugaredLogger
}

func NewZapLogger(l *zap.SugaredLogger) *ZapLogger {
	return &ZapLogger{
		log: l,
	}
}

func (l *ZapLogger) Named(name string) Logger {
	return &ZapLogger{
		log: l.log.Named(name),
	}
}

func (l *ZapLogger) Info(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.log.Infow(msg, keysAndValues...)
}

func (l *ZapLogger) Warn(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.log.Warnw(msg, keysAndValues...)
}

func (l *ZapLogger) Error(ctx context.Context, msg string, err error, keysAndValues ...interface{}) {
	l.log.Errorw(msg, append([]interface{}{"reason", err}, keysAndValues...)...)
}

func (l *ZapLogger) Panic(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.log.Panicw(msg, keysAndValues...)
}

func (l *ZapLogger) Fatal(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.log.Fatalw(msg, keysAndValues...)
}

func (l *ZapLogger) Debug(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.log.Debugw(msg, keysAndValues...)
}
