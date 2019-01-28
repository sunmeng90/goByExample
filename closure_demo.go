package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()
	fmt.Println("-----nextInt-----")
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextIntV2 := intSeq()
	fmt.Println("-----nextIntV2-----")
	fmt.Println(nextIntV2())
	fmt.Println(nextIntV2())
	fmt.Println(nextIntV2())
	fmt.Println(nextIntV2())
}
