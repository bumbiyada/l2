package main

import (
	"fmt"
	"time"
)

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		fmt.Println("starting chan")
		go func() {
			defer close(c)
			fmt.Println("starting goroutine")
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	<-aboba(
		sig(2*time.Second),
		sig(5*time.Second),
		sig(1*time.Second),
		sig(10*time.Second),
		sig(8*time.Second),
	)
	fmt.Printf("phone after %v", time.Since(start))
}

// our function that listens all channels and return
func aboba(channels ...<-chan interface{}) <-chan interface{} {
	fmt.Println("returning 0 chan")
	for {
		for idx, ch := range channels {
			select {
			case _, ok := <-ch:
				if !ok {
					fmt.Println("chan is closed", idx)
					// IDK WHAT TO RETURN
					return channels[idx]
				}
			default:
				fmt.Println("NOT CLOSED", idx)
				time.Sleep(time.Microsecond * 100)
			}
		}
	}
}
