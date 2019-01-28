package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("------------array------------")
	array1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array1)
	fmt.Println("array1 at:", unsafe.Pointer(&array1))
	fmt.Println("array1_0 at:", unsafe.Pointer(&array1[0]))
	fmt.Println("array1_1 at:", unsafe.Pointer(&array1[1]))

	fmt.Println("------------slice------------")
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)
	fmt.Println("array1 at:", &slice1, ", len: ", len(slice1), ", cat: ", cap(slice1))
	fmt.Println("slice1 at:", unsafe.Pointer(&slice1))
	fmt.Println("slice1_0 at:", unsafe.Pointer(&slice1[0]))
	fmt.Println("slice1_1 at:", unsafe.Pointer(&slice1[1]))

	fmt.Println("------------slice from slice------------")
	slice1_a := slice1[:3]
	fmt.Println(slice1_a)
	fmt.Println("slice1_a at:", &slice1_a, ", len: ", len(slice1_a), ", cat: ", cap(slice1_a))
	fmt.Println("slice1_a at:", unsafe.Pointer(&slice1_a))
	fmt.Println("slice1_a_0 at:", unsafe.Pointer(&slice1_a[0]))
	fmt.Println("slice1_a_1 at:", unsafe.Pointer(&slice1_a[1]))
	slice1_a[0] = 111
	fmt.Println("------------slice from slice modified(slice1_a)------------")
	fmt.Println("slice1_a at:", &slice1_a)
	fmt.Println("slice1_a at:", unsafe.Pointer(&slice1_a))
	fmt.Println("slice1_a_0 after modified at:", unsafe.Pointer(&slice1_a[0]))
	fmt.Println("------------slice modified(slice1)------------")
	fmt.Println(slice1)
	fmt.Println("slice1 at:", &slice1)
	fmt.Println("slice1 at:", unsafe.Pointer(&slice1))
	fmt.Println("slice1_0 at:", unsafe.Pointer(&slice1[0]))

	slice1_appended := append(slice1, 999)
	fmt.Println("------------slice(slice1_appended)------------")
	fmt.Println("slice1_appended at:", &slice1_appended, ", len: ", len(slice1_appended), ", cat: ", cap(slice1_appended))
	fmt.Println("slice1_appended at:", &slice1_appended)
	fmt.Println("slice1_appended at:", unsafe.Pointer(&slice1_appended))
	fmt.Println("slice1_appended_0 after modified at:", unsafe.Pointer(&slice1_appended[0]))
	fmt.Println("------------slice1_a(slice1_appended)------------")
	fmt.Println(slice1_a)
	fmt.Println("slice1_a at:", &slice1_a)
	fmt.Println("slice1_a at:", unsafe.Pointer(&slice1_a))
	fmt.Println("slice1_a at:", unsafe.Pointer(&slice1_a[0]))

	slice1[0] = 111111
	fmt.Println("------------slice append modified(slice1_appended)------------")
	fmt.Println("slice1_a at:", &slice1_appended)
	fmt.Println("slice1_a at:", unsafe.Pointer(&slice1_appended))
	fmt.Println("slice1_a_0 after modified at:", unsafe.Pointer(&slice1_appended[0]))
	fmt.Println("------------slice modified(slice1)------------")
	fmt.Println(slice1)
	fmt.Println("slice1 at:", &slice1)
	fmt.Println("slice1 at:", unsafe.Pointer(&slice1))
	fmt.Println("slice1_0 at:", unsafe.Pointer(&slice1[0]))

}
