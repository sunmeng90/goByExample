package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// A goroutine holds a state(map), which can be access by other goroutines by sending read operation or write operation through different channel.
// different operation are wrapped in to different struct and send through typed channels

// read operation
// get a value from state by key and value will be sent back through resp channel
type readOp struct {
	key  int
	resp chan int
}

// write operation
// write a key-val pair to the statue and true/false will be send back through resp channel to indicate the success of write
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	var readOps uint64
	var writeOps uint64

	reads := make(chan *readOp)   // read channel that read operation will be send through
	writes := make(chan *writeOp) // write channel that write operation will be send through

	//state holder goroutine
	go func() {
		var state = make(map[int]int)
		//handling readOp/writeOp
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()
	//create goroutine that will send read operations
	for r := 0; r < 100; r++ {
		go func() {
			//keeping sending read operation
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				//count the read operations finished(atomic change)
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	//create goroutine that will send write operations
	for w := 0; w < 10; w++ {
		go func() {
			//keeping sending write operation
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
