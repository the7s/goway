package goroutine

import (
	"fmt"
	"math/rand"
)

func SelectMain() {
	ch1 := make(chan string)
	ch2 := make(chan int)

	go func(ch1 chan string) {
		ch1 <- "test1"
	}(ch1)

	go func(ch2 chan int) {
		ch2 <- rand.Int()
	}(ch2)

	for {
		select {
		case s1 := <-ch1:
			fmt.Println("s1=", s1)
		case s2 := <-ch2:
			fmt.Println("s2=", s2)
		default:
			break
		}
	}
}
