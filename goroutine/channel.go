package goroutine

import "fmt"

// Channel1 优雅的从通道中循环取值
func Channel1() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {

		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for {
			i, ok := <-ch1

			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()

	for i := range ch2 {
		fmt.Println("i value is : ", i)
	}

}

// UniDirectChannel 单向通道
func UniDirectChannel() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func(out chan<- int) {
		for i := 0; i < 100; i++ {
			out <- i
		}
		close(out)
	}(ch1)

	go func(in <-chan int, out chan<- int) {

		for {
			i, ok := <-in
			if !ok {
				break
			}
			out <- i * i
		}
		close(out)
	}(ch1, ch2)

	func(out <-chan int) {
		for i := range out {
			fmt.Println("i is :", i)
		}
	}(ch2)
}
