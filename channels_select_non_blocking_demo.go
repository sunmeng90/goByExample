package main

import (
	"fmt"
)

//select用于在多个channel上同时进行侦听并收发消息，当任何一个case满足条件时即执行，如果没有可执行的case则会执行default的case，如果没有指定defaultcase，则会阻塞程序，select的语法：
//
//select {
//
//   case communication clause :
//
//      statement(s);
//
//   case communication clause :
//
//      statement(s);
//
//   /*可以定义任意数量的 case */
//
//   default : /*可选 */
//
//      statement(s);
//
//}
// case default in select will be executed immediately when no values arrives on any of channel in case branches
//
// Basic sends and receives on channels are blocking. However, we can use select with a default clause to implement non-blocking sends, receives, and even non-blocking multi-way selects.
func main() {
	messages := make(chan string)
	signals := make(chan string)

	// Here’s a non-blocking receive. If a value is available on messages then select will take the <-messages case with that value.
	// If not it will immediately take the default case.
	select {
	case msg := <-messages:
		fmt.Println("Received message", msg)
	default: //not blocking
		fmt.Println("No message received")
	}

	msg := "hi"
	// A non-blocking send works similarly. Here msg cannot be sent to the messages channel,
	// because the channel has no buffer and there is no receiver. Therefore the default case is selected.
	select {
	case messages <- msg:
		fmt.Println("Sent message", msg)
	default:
		fmt.Println("No message sent")
	}

	// We can use multiple cases above the default clause to implement a multi-way non-blocking select.
	// Here we attempt non-blocking receives on both messages and signals.
	select {
	case msg := <-messages:
		fmt.Println("Received message", msg)
	case sig := <-signals:
		fmt.Println("Received signal", sig)
	default:
		fmt.Println("No activity")

	}
}
