package main

import "fmt"

// 空接口可以传入任意类型的对象

type A interface {
	//什么都不定义
}

type Person1 struct {
	name string
	age  int
}

func main() {
	var a1 A = Person1{ //注意这是拿空接口A定义的a1
		name: "renyujie",
		age:  22,
	}
	//任意类型

	var a2 A = 10
	var a3 A = "hahaha"
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)

	//接收任意类型的数据
	test1(a1)
	test1(3.14)
	test1("sassa")
	//匿名空接口
	test2(a1)
	test2(3.12)
	test2("Sasa")

	//空接口还可以作为一个容器，储存任意类型的数据，比如对map
	map1 := make(map[string]interface{})
	map1["name"] = "renyujie"
	map1["age"] = 22
	map1["friend"] = Person1{
		name: "yuwenge",
		age:  23,
	}
	fmt.Println(map1)

	slice1 := make([]interface{}, 0, 10) //任意"类型"，所以类型写成interface{},又因为是切片，所以type:[]interface{}
	slice1 = append(slice1, a1, 3124, map1, "Ssssssss")
	fmt.Println(slice1)
	fmt.Printf("容器的类型是%T\n", slice1)

	//测试把上面这个接收了任何类型的切片当做函数的输入的参数，同时，保证函数定义的时候传参也带有空接口
	text3(slice1)

}

//A是一个空接口，a可以被看做任意类型的数据
func test1(a A) {
	fmt.Println("我不管什么类型我都可以接收", a)

}

//匿名空接口，省去了定义一个type A interface的麻烦
func test2(a interface{}) {
	fmt.Println("匿名空接口也什么都可以可以接收", a)
}

//设计一个函数，接收slice为参数 可以接收任意类型的数据

func text3(slice []interface{}) {
	for i := 0; i < len(slice); i++ {
		fmt.Printf("第%d个数据：%v\n", i+1, slice[i])
	}
}
