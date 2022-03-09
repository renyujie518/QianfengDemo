package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//复制文件
func main() {
	srcFile := "/Users/renyujie/go/src/QianfengDemo/homework/IoTest/a.txt"
	destFile := "/Users/renyujie/go/src/QianfengDemo/homework/IoTest/b.txt"
	//total,err := Copyfile1(srcFile,destFile)
	//total,err := Copyfile2(srcFile,destFile)
	total, err := Copyfile3(srcFile, destFile)
	Checkerr(err)
	fmt.Println("拷贝的字节数", total)

}

func Copyfile3(srcFile string, destFile string) (int, error) {
	bs, err := ioutil.ReadFile(srcFile)
	if err != nil { //ReadFile终止条件是bil,不是io.EOF
		return 0, err
	}
	err = ioutil.WriteFile(destFile, bs, os.ModePerm)
	if err != nil {
		return 0, err
	}
	return len(bs), nil

}
func Copyfile2(srcFile string, destFile string) (int64, error) {
	file1, err := os.Open(srcFile) //源文件只读即可，所以用open
	Checkerr(err)
	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	Checkerr(err)
	defer file1.Close()
	defer file2.Close()
	return io.Copy(file2, file1) //注意，这里是先目标后源文件
}

func Copyfile1(srcFile string, destFile string) (int, error) {
	file1, err := os.Open(srcFile) //源文件只读即可，所以用open
	Checkerr(err)
	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	Checkerr(err)
	defer file1.Close()
	defer file2.Close()
	//读写
	bs := make([]byte, 1024, 1024)
	n := -1 //读取得我数据个数
	total := 0
	for {
		n, err = file1.Read(bs)
		if err == io.EOF || n == 0 {
			fmt.Println("拷贝完毕")
			break

		} else if err != nil {
			fmt.Println("拷贝出现错误")
			return total, err
		}
		total += n
		file2.Write(bs[:n])
	}
	return total, nil

}

func Checkerr(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
