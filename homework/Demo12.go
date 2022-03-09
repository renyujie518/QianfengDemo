package main

import "fmt"

/*
利用结构体的嵌套模拟java中的继承关系
1 继承性(is-a )   A中的字段对于B来说是提升字段，可以直接访问
type A struct{
	xxx
}
type B struct{
	A
}

2 聚合性（has a）
type A struct{
   xxx
}
type B struct {
	a A
}
*/
func main() {
	p1 := Person{
		name: "renyujie",
		age:  22,
	}
	fmt.Println(p1.name, p1.age) //直接访问父类对象和方法
	p1.eat()

	s1 := Student{
		Person: Person{name: "yuwenge", age: 22},
		school: "华电",
	}
	fmt.Println()
	s1.eat()                                //子类访问父类的方法！！！
	fmt.Println(s1.name, s1.age, s1.school) //实际上是s1.Person.name,但由于匿名了，字段直接就是名字
	s1.study()                              //子类访问针对自己的方法
	s1.eat()                                //子类重写父类的方法，此时第36行的执行结果也变了，"覆盖"重写

}

//定义一个父类
type Person struct {
	name string
	age  int
}

//定义一个子类
type Student struct {
	Person        //匿名字段 模拟继承性
	school string //子类自己独有的字段
}

func (p Person) eat() {
	fmt.Print("父类的方法：所有人可以吃饭")
}

func (s Student) study() {
	fmt.Println("子类新增的方法：学生可以学习")
}

func (s Student) eat() {
	fmt.Println("子类重写父类的方法：学生在学校吃饭")

}
