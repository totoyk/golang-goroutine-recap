package main

import "time"

func sayHello() {
	for i := 0; i < 5; i++ {
		println("Hello, World!")
		time.Sleep(100 * time.Millisecond)
	}
}

func channelExample() {
	ch := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- "Hello from channel!"
			time.Sleep(100 * time.Millisecond)
		}
		close(ch)
	}()

	for msg := range ch {
		println(msg)
	}
}

func main() {
	go sayHello()
	for i := 0; i < 5; i++ {
		println("Main function is running")
		time.Sleep(100 * time.Millisecond)
	}

	channelExample()
}
