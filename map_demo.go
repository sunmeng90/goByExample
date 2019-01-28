package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["key0"] = 0
	m["key1"] = 1
	m["key2"] = 2
	fmt.Println(m)
	fmt.Println(m["key0"])
	fmt.Println(m["keyzero"])
	key0_val, key0_exists := m["key0"] //second value indicate whether the key exists
	keyZero_val, keyZero_exists := m["keyzero"]
	fmt.Println("key0 value: ", key0_val, ", exists: ", key0_exists)
	fmt.Println("keyzero value: ", keyZero_val, ", exists: ", keyZero_exists)

	persons := make(map[string]int)
	persons["张三"] = 19
	mp := &persons
	fmt.Printf("原始map的内存地址是：%p\n", mp)
	modify(persons)
	fmt.Println("map值被修改了，新值为:", persons)
}
func modify(p map[string]int) {
	fmt.Printf("函数里接收到map的内存地址是：%p\n", &p) //TODO: why pointer is different
	p["张三"] = 20
}
