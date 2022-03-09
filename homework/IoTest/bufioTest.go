package main

import (
	"bufio"
	"fmt"
	"os"
)

//带缓存的读写
func main() {

	filename := "/Users/renyujie/go/src/QianfengDemo/homework/IoTest/a.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	//创建Reader对象  用NewReaderd的read是一种带缓存的高效读法
	b1 := bufio.NewReader(file)
	p := make([]byte, 1024) //默认的NewReader的size是4096，p超过这个缓冲区则没用
	n1, err := b1.Read(p)
	fmt.Println(n1) //读的字节的数量
	fmt.Println(string(p[:n1]))

	//Readline() 读一行  不建议使用
	//data,flag,err := b1.ReadLine()
	//if err!= nil{
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(flag)
	//fmt.Println(string(data))

	//直接读成字符串 ReadString
	//s1,err := b1.ReadString('\n')//读到换行为止，直接返回字符串
	//if err!= nil{
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(s1)
	//fmt.Println("在打印一行")
	//s1,err = b1.ReadString('\n')//读到换行为止，直接返回字符串
	//if err!= nil{
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(s1)

	//for  {
	//	s1,err := b1.ReadString('j')
	//	if err == io.EOF {
	//		fmt.Println("读取完毕++++++++++")
	//	}
	//	fmt.Println(s1)
	//
	//}

}
