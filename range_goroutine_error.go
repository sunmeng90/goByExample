package main

import "fmt"

func main() {
	sliceA := [...]string{"hello", "world"}
	for _, item := range sliceA {
		fmt.Printf("[value] %v\n", item) //hello world
		go func() {
			//!!!wrong don't do this
			fmt.Printf("[wrong]" + item + "\n") //world world
		}()
	}

	for _, item := range sliceA {
		fmt.Printf("[value] %v\n", item) //hello world
		go func(str string) {
			fmt.Printf("[correct]" + str + "\n") //hello world
		}(item)
	}

	fmt.Scanln()
}
