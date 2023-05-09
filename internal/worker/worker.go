package worker

import (
	"context"
	"log"
	"time"

	"go.uber.org/zap"
)

func Worker(ctx context.Context, ticker <-chan time.Time, fn func() error) {
	defer func() { log.Printf("worker stopped") }()

	for {
		select {
		case <-ticker:
			if err := fn(); err != nil {
				zap.S().Error(err)
			}
		case <-ctx.Done():
			return
		}
	}
}
