package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go task(ctx, "task1")
	go task(ctx, "task2")
	go task(ctx, "task3")

	time.Sleep(10 * time.Second)
	fmt.Println("time's up, stop task")
	cancel()
	time.Sleep(5 * time.Second)
}

func task(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, ",stopped")
			return
		default:
			fmt.Println(name, "doing.")
			time.Sleep(2 * time.Second)
		}
	}
}
