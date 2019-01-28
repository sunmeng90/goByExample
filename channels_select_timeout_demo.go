package main

import (
	"fmt"
	"time"
)

//select will be executed once, whenever a value arrives from a channel, so we can set a channel that will send out a value after a duration to simulate a timeout
func main() {
	chann1 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		chann1 <- "go_a"
	}()

	//value from chann1 will arrives before vale from channel(return from time.After)
	select {
	case msg1 := <-chann1:
		fmt.Println("chann1: ", msg1)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout 1")
	}

	chann2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		chann2 <- "go_b"
	}()

	//value from channel(return from time.After) will arrives before value from chann2
	select {
	case msg2 := <-chann2:
		fmt.Println("chann2: ", msg2)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 2")
	}
}
