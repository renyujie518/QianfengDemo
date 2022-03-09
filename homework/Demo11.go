package main

import "fmt"

func main() {
	w1 := Worker{
		name: "小花",
		age:  10,
		sex:  "female",
	}
	w1.work()
	w2 := Worker{
		name: "小明",
		age:  22,
		sex:  "male",
	}
	w2.realwork()
	w1.realwork() //相同的方法名可以作用在不同的对象上

}

//首先定义一个工人结构体
type Worker struct {
	name string
	age  int
	sex  string
}

//定义方法行为   func (接受者) 方法名（参数）（返回值）{}
func (w Worker) work() {
	fmt.Println(w.name, "在工作")
}

//接受者为指针类型的方法，这类方法一旦调用就会改变接受者的值
func (p *Worker) realwork() {
	fmt.Println(p.name, "在真正工作") //实际上应该是*p.name,这里是一种省略写法
}
