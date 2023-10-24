package main

import (
	"context"
	"fmt"
	"time"
)

type ContextValueKey string

func main() {
	ctx := context.Background()
	var v ContextValueKey = "Galdamez"
	var k ContextValueKey = "my-key"
	ctx = context.WithValue(ctx, k, v)

	viewContext(ctx)

	ctx2, cancelF := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancelF()

	myProcess(ctx2)

}

func viewContext(ctx context.Context) {
	var k ContextValueKey = "my-key"
	fmt.Printf("my value is %s\n", ctx.Value(k))
}

func myProcess(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ooh, time out")
			return
		default:
			fmt.Println("doing some process")
		}

		time.Sleep(200 * time.Millisecond)
	}
}
