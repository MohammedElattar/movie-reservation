package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	wg.Go(func() {
		if err := printGreeting(ctx); err != nil {
			fmt.Printf("cannot print greeting %v\n", err)
			cancel()
		}
	})

	wg.Go(func() {
		if err := printBye(ctx); err != nil {
			fmt.Printf("cannot print bye %v\n", err)
		}
	})

	wg.Wait()
}

func printGreeting(ctx context.Context) error {
	greeting, err := genGreeting(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("%s world!\n", greeting)

	return nil
}

func printBye(ctx context.Context) error {
	bye, err := genBye(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("good %s\n", bye)

	return nil
}

func genGreeting(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	locale, err := locale(ctx)
	if err != nil {
		return "", err
	}

	switch locale {
	case "en":
		return "hello", nil
	default:
	}
	return "", fmt.Errorf("unsupported locale")
}

func genBye(ctx context.Context) (string, error) {
	locale, err := locale(ctx)
	if err != nil {
		return "", err
	}

	switch locale {
	case "en":
		return "bye", nil
	default:
		return "", fmt.Errorf("unsupported locale")
	}
}

func locale(ctx context.Context) (string, error) {
	if deadline, ok := ctx.Deadline(); ok {
		if deadline.Sub(time.Now().Add(1*time.Second)) <= 0 {
			return "", context.DeadlineExceeded
		}
	}

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(1 * time.Minute):
	}

	return "en", nil
}
