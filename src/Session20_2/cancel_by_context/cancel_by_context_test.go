package cancel_by_context_test

import (
	"fmt"
	"time"
	"testing"
	"context"
)


func isCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	for i:= 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancelled(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "cancelled")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)
}