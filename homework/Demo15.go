package main

import (
	"fmt"
	"math"
)

type Shape interface {
	zhouchang() float64
	mianji() float64 //这个float64是返回值类型
}

//三角形
type Triangle struct {
	a, b, c float64
}

func (t Triangle) zhouchang() float64 {
	return t.a + t.b + t.c
}
func (t Triangle) mianji() float64 {
	p := t.zhouchang() / 2 //由于先实现过zhouchang的方法，既然都是三角形对于Shape的实现类，说明是可以相互调用的
	s := math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
	return s
}

//圆形
type Circle struct {
	r float64
}

func (c Circle) zhouchang() float64 {
	return 2 * c.r * math.Pi
}
func (c Circle) mianji() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

//断言实现到底传入接口的是哪一种子对象
func getType(s Shape) {
	//断言
	if ins, ok := s.(Triangle); ok { //从括号内看，s是不是三角形，ok的话,返回实例ins同时ok为true
		fmt.Println("是三角形，三边是", ins.a, ins.b, ins.c)
	} else if ins, ok := s.(Circle); ok {
		fmt.Println("是圆形，半径是", ins.r)
	} else {
		fmt.Println("什么都不是")
	}
}

func main() {
	var t1 Triangle = Triangle{
		a: 3,
		b: 4,
		c: 5,
	}
	var c1 Circle = Circle{3}

	fmt.Println("三角形周长： ", t1.zhouchang())
	fmt.Println("三角形面积： ", t1.mianji())
	fmt.Println("圆形面积： ", c1.zhouchang())
	fmt.Println("圆形面积： ", c1.mianji())

	//声明一个借口类型
	var s1 Shape
	s1 = t1
	fmt.Println(s1.zhouchang())
	fmt.Println(s1.mianji()) //接口赋值了三角形，就只能调用三角形实现的方法，同时也不能访问t.a这些特有字段
	var s2 Shape
	s2 = c1

	//断言测试
	getType(s1)
	getType(s2)
	getType(t1)
	getType(c1)

}
