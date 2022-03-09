package main

import (
	"fmt"
	"reflect"
	"strconv"
)

//用来检测储存在接口变量内部的pair对中(value；concrete type)的一种机制，运行时可以动态的获取这两个变量来了解自己
//反射是针对接口来说的，通反射得知接口的动态信息
/*
反射的三大功能
1.从接口的pair里得到类型和值，TypeOf(），ValueOf()
2.与第一条相反，通过将ValueOf得到的返回值在通过包下的函数Interface（）转成接口变量（真实内容）
3.反射变量的本质是储存了原变量本身，所以对反射变量的操作也会改变原变量,这时候需要传递一个地址值
*/

type Person17 struct {
	Name string
	Age  int
	Sex  string
}

func (p Person17) Say(msg string) {
	fmt.Println("hello", msg)
}

func (p Person17) Personinfo() {
	fmt.Printf("姓名：%s,年龄：%d,性别:，%s", p.Name, p.Age, p.Sex)
}

func (p Person17) Test(i, j int, s string) {
	fmt.Println(i, j, s)
}

func Getmessage(input interface{}) {
	getType := reflect.TypeOf(input)
	fmt.Println("get type is ", getType.Name())
	fmt.Println("get kind is ", getType.Kind())
	getValue := reflect.ValueOf(input)
	fmt.Println("get allFiled is ", getValue)

	/*
		获取真实内容的思路
		1.获取type对象
		2，通过Filed()获取每个字段
		3、转换为Interface()
	*/
	for i := 0; i < getType.NumField(); i++ {
		filed := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("字段名称 %s,字段类型 %s,字段的数值 %v\n", filed.Name, filed.Type, value)
	}
	//获取方法明细
	for i := 0; i < getType.NumMethod(); i++ {
		method := getType.Method(i)
		fmt.Printf("方法名称 %s,方法类型 %v\n", method.Name, method.Type)
	}

}

func main() {
	var x float64 = 4.3
	fmt.Println("type", reflect.TypeOf(x))
	fmt.Println("value", reflect.ValueOf(x))

	value := reflect.ValueOf(x)
	convertValue := value.Interface().(float64)
	println(convertValue)

	Pointvalue := reflect.ValueOf(&x)
	convertPointer := Pointvalue.Interface().(*float64)
	println(convertPointer)
	newvalue := Pointvalue.Elem() //获取指针所指的那个Value对象（获取原始值对应的反射对象）
	fmt.Println("类型", newvalue.Type())
	fmt.Println("是否可以修改数据", newvalue.CanSet())
	//重新赋值，对应用法3
	newvalue.SetFloat(22212.21)
	fmt.Println("修改后的值", x)

	fmt.Println("====================")

	p1 := Person17{
		Name: "张三",
		Age:  10,
		Sex:  "男",
	}
	Getmessage(p1)

	//通过反射区更改对象的数值
	valuePerson := reflect.ValueOf(&p1) //结构体是值传递，所以传指针才能改变
	if valuePerson.Kind() == reflect.Ptr {
		newvaluePerson := valuePerson.Elem()
		fmt.Println("是否可以修改结构体数据", newvaluePerson.CanSet())

		valueFieldbyname := newvaluePerson.FieldByName("Name")
		valueFieldbyname.SetString("李四")
		valueFieldbyage := newvaluePerson.FieldByName("Age")
		valueFieldbyage.SetInt(22)
		fmt.Println("修改后的结构体:")
		p1.Personinfo()
		fmt.Println("========================")
		//通过反射进行方法的调用（高级用法）
		/*
			1.接口变量--》反射对象value
			2.获得方法对象 MethodByName()或者下标获取
			3.调用 call()
		*/
		p2 := Person17{
			Name: "千峰教育",
			Age:  100,
			Sex:  "女",
		}
		valuePerson2 := reflect.ValueOf(p2)
		fmt.Printf("kind: %s,type: %v\n", valuePerson2.Kind(), valuePerson2.Type())
		methodValue1 := valuePerson2.MethodByName("Personinfo")
		fmt.Printf("kind: %s,type: %v\n", methodValue1.Kind(), methodValue1.Type())
		methodValue2 := valuePerson2.MethodByName("Say")
		fmt.Printf("kind: %s,type: %v\n", methodValue2.Kind(), methodValue2.Type())
		methodValue3 := valuePerson2.MethodByName("Test")
		fmt.Printf("kind: %s,type: %v\n", methodValue3.Kind(), methodValue3.Type())

		//没参的方法传nil或者一个空切片到call
		args1 := make([]reflect.Value, 0)
		methodValue1.Call(args1)
		//有参数的要传参，多个参数顺序要严格对应，会自动依次放入call
		args2 := []reflect.Value{reflect.ValueOf("Valueof是把任意输入转为反射对象")}
		//注意这里不像给方法传参一样直接传字符串了，要把这个字符串转成反射对象才可以放入[]reflect.Value的切片
		fmt.Println()
		methodValue2.Call(args2)
		args3 := []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2), reflect.ValueOf("ssd")}
		methodValue3.Call(args3)
		fmt.Println("=============================")

		//通过反射进行函数的调用（高级用法），与方法类似
		funValue1 := reflect.ValueOf(fun1) //注意 其实函数也可以看做接口变量类型，而ValueOf(interface{})的输入是一个空接口，所以可以直接输入
		fmt.Printf("kind: %s,type: %v\n", funValue1.Kind(), funValue1.Type())
		funValue2 := reflect.ValueOf(fun2)
		fmt.Printf("kind: %s,type: %v\n", funValue2.Kind(), funValue2.Type())
		funValue3 := reflect.ValueOf(fun3)
		fmt.Printf("kind: %s,type: %v\n", funValue3.Kind(), funValue3.Type())
		funValue1.Call(nil)
		funValue2.Call([]reflect.Value{reflect.ValueOf(100), reflect.ValueOf("元")})
		//Call的返回值是一个Value[]
		resultValue := funValue3.Call([]reflect.Value{reflect.ValueOf(300), reflect.ValueOf("yuan")})
		fmt.Printf("返回的结果的类型是 %T\n", resultValue)
		fmt.Printf("kind: %s,type: %v\n", resultValue[0].Kind(), resultValue[0].Type())
		//可以在通过用法二还原,相当于调用这个函数
		s := resultValue[0].Interface().(string)
		fmt.Println(s)

	}
}
func fun1() {
	fmt.Println("我是函数1，无参的")
}
func fun2(i int, s string) {
	fmt.Printf("我是函数2，有参的，参数为%d,%s\n", i, s)
}
func fun3(i int, s string) string {
	fmt.Printf("我是函数3，有参的，还有返回值，参数为%d,%s\n", i, s)
	return s + strconv.Itoa(i)
}
