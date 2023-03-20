package garbagecollector

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/lvlBA/online_shop/internal/passport/db"
	"github.com/lvlBA/online_shop/pkg/logger"
)

type Token struct {
	log     logger.Logger
	expired time.Duration
	db      db.Service
	timeout time.Duration
}

type Config struct {
	Log     logger.Logger
	Expired time.Duration
	DB      db.Service
	Timeout time.Duration
}

func NewToken(cfg *Config) *Token {
	return &Token{
		log:     cfg.Log,
		expired: cfg.Expired,
		db:      cfg.DB,
		timeout: cfg.Timeout,
	}
}

func (t *Token) Observe(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(t.timeout)

	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			return
		case <-ticker.C:
			if err := t.db.Auth().DeleteOldTokens(ctx); err != nil {
				if !errors.Is(err, db.ErrorNotFound) {
					t.log.Error(ctx, "failed to delete old tokens", err)
				}
			}
		}
	}
}
