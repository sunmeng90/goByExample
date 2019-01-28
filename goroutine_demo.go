package main

import (
	"fmt"
	"time"
)

func counter(n int, prefix string) {
	for i := 0; i < n; i++ {
		fmt.Println(prefix+":", i)
		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	counter(5, "Counter-A")
	go counter(5, "Counter-B") //normal function can be run asynchronously using goroutine

	go func(msg string) {
		fmt.Println("Echo", msg)
	}("Hello") // echo message can possibly be displayed in the middle of Counter-B's counting

	fmt.Scanln() //wait for manual input to exit from main goroutine, otherwise main goroutine may exit before the sub-goroutines
	fmt.Println("done")

}
