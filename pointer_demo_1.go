package main

/*
   1>.我们希望改变对象的成员；
   2>.等价于函数的第一个参数是指针；
*/

import (
	"fmt"
)

type Point struct {
	X, Y float64
}

func (p *Point) ScaleBy(factor float64) { //想要修改p的值就得传指针类型"*Point"
	p.X *= factor // 等价于：X = X * factor，对对象P进行操作，修改其共有属性。
	p.Y *= factor
}

func main() {
	//两种调用方式：
	p := Point{100, 200}
	p.ScaleBy(2) //姿势一：直接调用
	fmt.Println(p)

	p1 := Point{100, 200} //姿势二：声明结构体后再用指针指向
	p2 := &p1             //使用结构体调用，再取其内存地址
	p2.ScaleBy(2)
	fmt.Println(p2)
	fmt.Println(*p2)
}
