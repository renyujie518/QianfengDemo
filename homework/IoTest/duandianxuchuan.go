package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
总的思路是边复制边记录下复制的总量
*/
func main() {
	srcFile := "/Users/renyujie/go/src/QianfengDemo/homework/IoTest/1.png"
	destFile := srcFile[strings.LastIndex(srcFile, "/")+1:]
	fmt.Println(destFile)
	//设置临时文件，记录复制到哪里了，整个文件确定复制完了会删除临时文件
	tempFile := destFile + "temp.txt"
	fmt.Println(tempFile)

	file1, err := os.Open(srcFile)
	Checkerrs(err)
	file2, err := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	Checkerrs(err)
	file3, err := os.OpenFile(tempFile, os.O_CREATE|os.O_RDWR, os.ModePerm)
	Checkerrs(err)

	defer file1.Close()
	defer file2.Close()
	//先读取临时文件的数据，再seek
	file3.Seek(0, io.SeekStart)
	bs := make([]byte, 100, 100)                   //bs装的是在临时文件中所记录的已经复制完的数据量是多少
	n1, _ := file3.Read(bs)                        //注意，在第一次读取临时文件的时候里面没任何数据，所以会报读错误，这里忽略这个错误即可
	countStr := string(bs[:n1])                    //这是临时文件里储存的已经传完的字节，先转string再转int
	count, _ := strconv.ParseInt(countStr, 10, 64) //转成10进制，最大64位的整数

	fmt.Println("上次传递的个数是", count)

	//设置新的读写位置
	file1.Seek(count, io.SeekStart)
	file2.Seek(count, io.SeekStart)
	data := make([]byte, 1024, 1024) //用来复制文件
	n2 := -1                         //读的数据量
	n3 := -1                         //写的数据量
	total := int(count)              //读取的总量，初始要设置成临时文件里的数量,因为有可能已经发生断点续传了，以后total再在count上累加

	//复制文件
	for {
		n2, err = file1.Read(data)
		if err == io.EOF || n2 == 0 {
			fmt.Println("复制完毕", total)
			file3.Close()
			os.Remove(tempFile)
			break
		}
		//写数据
		n3, err = file2.Write(data[:n2])
		total += n3

		//将总量存到临时文件，而且覆盖临时文件上次的数据（所以seek从0开始）
		file3.Seek(0, io.SeekStart)
		file3.WriteString(strconv.Itoa(total))
		fmt.Printf("已经复制的数量: %d\n", total)

		////模拟断电
		//if total > 8000{
		//	panic("假装断电了")
		//}
	}

}
func Checkerrs(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
