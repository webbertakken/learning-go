package main

import (
	"context"
	"fmt"
	"time"
)

func sampleOperation(ctx context.Context, msg string, delay time.Duration) <-chan string {
	out := make(chan string)

	go func() {
		for {
			select {
			case <-time.After(delay * time.Millisecond):
				out <- fmt.Sprintf("%v operation completed", msg)
				return
			case <-ctx.Done():
				out <- fmt.Sprintf("%v aborted", msg)
				return
			}
		}
	}()

	return out
}

func main() {
	ctx := context.Background()
	ctx, cancelCtx := context.WithCancel(ctx)

	webServer := sampleOperation(ctx, "webserver", 200)
	microservice := sampleOperation(ctx, "microservice", 500)
	database := sampleOperation(ctx, "database", 900)

MainLoop:
	for {
		select {
		case value := <-webServer:
			fmt.Println(value)
		case value := <-microservice:
			fmt.Println(value)
			fmt.Println("cancelling context")
			cancelCtx()
			break MainLoop
		case value := <-database:
			fmt.Println(value)
		}
	}

	fmt.Println(<-database)
}
