package main

import "fmt"

/*

Channels are tunnel to exchange message between two goroutines

Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.

Create a new channel with make(chan val-type). Channels are typed by the values they convey.
Send a value into a channel using the channel <- syntax. Here we send "ping" to the messages channel we made above, from a new goroutine.

The <-channel syntax receives a value from the channel. Here we’ll receive the "ping" message we sent above and print it out.

When we run the program the "ping" message is successfully passed from one goroutine to another via our channel.

By default sends and receives block until both the sender and receiver are ready. This property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization.

We can use channels to synchronize execution across goroutines
*/
func main() {
	chan_message := make(chan string)

	go func() {
		s_msg := "hello"
		fmt.Println("Send: ", s_msg)
		chan_message <- s_msg
	}()

	go func() {
		r_msg := <-chan_message
		fmt.Println("Receive: ", r_msg)
	}()

	fmt.Scanln()
}
