package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	c := context.Background()

	ctx, CancelFunc := context.WithCancel(c)

	defer CancelFunc()
	go func(ctx context.Context) {
		for {
			fmt.Println("123")
			time.Sleep(time.Second)
			select {
			case <-ctx.Done():
				fmt.Println("外部结束")
				return
			default:
			}
		}

	}(ctx)
	time.Sleep(time.Second * 3)
	CancelFunc()
	time.Sleep(time.Second * 10)

}
