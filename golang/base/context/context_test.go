package context1

import (
	"context"
	"testing"
	"time"
)

func isCanceled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCanceled(ctx) {
					break
				}

				time.Sleep(time.Millisecond * 5)
			}
			t.Log(i, "Canneled")
		}(i, ctx)
	}

	cancel()
	time.Sleep(time.Second)
}
