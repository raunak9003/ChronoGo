package main

import (
	"context"
	"time"
)

func operationWithTimeout(parentctx context.Context, duration time.Duration) (string, error) {
	ctx, cancel := context.WithTimeout(parentctx, duration)
	defer cancel()
	return doOperation(ctx)
}

func doOperation(ctx context.Context) (string, error) {
	resultCh := make(chan string, 1)
	errCh := make(chan error, 1)
	go func() {
		result, err := expensiveOperation()
		select {
		case <-ctx.Done():
			return
		case result <- result:
			if err != nil {
				errCh <- err
			}
		}
	}()
	// Await results or handle cancellation
	select {
	case result := <-resultCh:
		return result, <-errCh
	case <-ctx.Done():
		return "", ctx.Err() // Return the context error if canceled
	}
}
