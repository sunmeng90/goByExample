package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	for key, num := range nums {
		fmt.Println("element @", key, ", value: ", num)
	}

	for _, num := range nums {
		fmt.Println("element value: ", num)
	}

	kvs := map[string]string{"key0": "str0", "key1": "str1"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "hello" {
		fmt.Println(i, c)
	}
}
