package main

import (
	"fmt"
	"time"
)

//Go’s select lets you wait on multiple channel operations. Combining goroutines and channels with select is a powerful feature of Go.
func main() {
	chann1 := make(chan string)
	chann2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		chann1 <- "go_a"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		chann2 <- "go_b"
	}()

	//We’ll use select to await both of these values simultaneously, printing each one as it arrives.
	for i := 0; i < 2; i++ { //wait for two values
		select {
		case msg1 := <-chann1:
			fmt.Println("chann1: ", msg1)
		case msg2 := <-chann2:
			fmt.Println("chann2: ", msg2)
		}
	}
}
