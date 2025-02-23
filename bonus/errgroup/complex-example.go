package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func processTask(id int) error {
	delay := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(delay)

	if rand.Float32() < 0.2 {
		return fmt.Errorf("error processing task %d", id)
	}

	fmt.Printf("Task %d completed in %s\n", id, delay)
	return nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	maxConcurrency := 5
	semaphore := make(chan struct{}, maxConcurrency)

	var g errgroup.Group
	var mu sync.Mutex
	totalTasks := 20
	completedTasks := 0

	for i := 1; i <= totalTasks; i++ {
		i := i
		g.Go(func() error {
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				if err := processTask(i); err != nil {
					cancel()
					return err
				}

				mu.Lock()
				completedTasks++
				mu.Unlock()

				return nil
			}
		})
	}

	err := g.Wait()

}
