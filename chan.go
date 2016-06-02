// https://gobyexample.com/select

package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan interface{})
	c2 := make(chan interface{})

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- 2
	}()

	for range []chan interface{}{c1, c2} {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
