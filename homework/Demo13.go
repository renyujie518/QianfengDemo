package main

import "fmt"

//接口可以看做是抽象类的集合
//接口最大的作用就是解耦合 功能的定义和实现分开，接口只是定义方法或者说功能有哪些
//谁实现了这些方法就是实现类
func main() {
	//测试1
	fmt.Println("*****接口对象访问自己的变量")
	m1 := Mouse{name: "罗技"}
	k1 := keyboard{name: "cherry"}
	fmt.Println(m1.name)
	fmt.Println(k1.name)
	//测试2
	fmt.Println("******接口对象调用抽象方法，也是自己覆盖重写过的")
	m1.start()
	m1.end()
	k1.start()
	k1.end()
	//测试3
	fmt.Println("*****既然实现过USB类的方法，就是接口的实现，就可以当做USB对象传入")
	textInterface(m1)
	textInterface(k1)
	//测试4
	fmt.Println("实现类可以直接赋值接口对象,但接口对象不能访问实现对象内独有的变量和方法，比如name")
	var usb1 USB
	usb1 = m1  //实现类可以直接赋值接口对象
	usb1.end() //这个usb1就可以看做一个鼠标
	usb1.start()
	//usb.name   错误写法 接口对象不能访问实现对象内独有的变量和方法。注意，这将引起多态

	/*
	   多态实际上针对子类的，因为直接定义父类的对象，就如上所说，这个父类对象就没法访问子类独有的方法和变量
	   好像只能单独在定义子类对象才行，子类是特殊的父类对象，既可以访问自己的，也可以完全访问父类的
	   这种定义的类型不同导致能访问的"权限"的不同，这就叫多态
	*/
	fmt.Println("不是所有的usb设备都是键盘，所以不能访问键盘新增的方法insert")
	var usb2 USB
	usb2 = k1 //这个usb2就可以看做一个键盘,但k1本身定义的时候是键盘，现在被赋值给usb,这就是多态
	// 不是所有的usb设备都是键盘，所以不能访问键盘新增的方法insert
	usb2.start()
	usb2.end()

	//接口的用法：
	//传入接口的任意实现类对象都可以作为参数
	//！！！！！接口只要被实现了，所实现的对象都可以给接口赋值（给骨骼填血肉）（非侵入式）
	var arr [3]USB
	arr[0] = m1
	arr[1] = k1
	fmt.Println(arr)

}

type USB interface {
	start()
	end()
}

//实现类(隐性的实现，即直接实现即可)
type Mouse struct {
	name string
}

type keyboard struct {
	name string
}

//鼠标实现接口USBz中 的方法
func (m Mouse) start() {
	fmt.Println("鼠标插入")

}
func (m Mouse) end() {
	fmt.Println("鼠标拔出")

}

//键盘实现接口USB中 的方法
func (k keyboard) start() {
	fmt.Println("键盘打字")

}
func (k keyboard) end() {
	fmt.Println("键盘拔出")

}

//测试方法
func textInterface(usb USB) {
	usb.start()
	usb.end()
}

//键盘新增方法
func (k keyboard) insert() {
	fmt.Println("键盘可以打字")
}
