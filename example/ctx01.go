package example

import (
	"context"
	"fmt"
	"time"
)

// WithCancel
func Ctx01() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("Done", ctx.Err())
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 10 {
			break
		}
	}
}

func Ctx02() {
	d := time.Now().Add(time.Second * 1)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-time.After(time.Second * 2):
		fmt.Println("overslept")
	}
}
