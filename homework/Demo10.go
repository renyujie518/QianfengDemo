package main

import (
	"fmt"
	"os"
	"time"
)

//自定义error，针对打开文件路径不存在这个错误设置自己的比较全面的Error() string
//首先创建一个struct类型来表示错误
type PathError struct {
	path     string
	op       string
	creatime string
	message  string
}

//实现error的接口
func (p *PathError) Error() string {
	//打印上面丰富的类型的错误即可
	//return fmt.Sprint("path = %s\nop = %s\ncreatime =%s\nmessage = %s ",p.path,
	//	p.op,p.creatime,p.message)
	return fmt.Sprintf("\npath=%s \nop=%s \ncreateTime=%s \nmessage=%s", p.path,
		p.op, p.creatime, p.message)
}

//实现打开文件和错误返回
func Open(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return &PathError{
			path:     filename,
			op:       "read",
			message:  err.Error(), //实现类的调用，用作显示
			creatime: fmt.Sprint("%v", time.Now()),
		}
	}
	defer file.Close() //延迟close
	return nil         //默认没有错
}

func main() {
	//调用open函数
	err := Open("/Users/renyujie/go/src/QianfengDemo/homework/Demo001.go")
	//设置断言
	switch v := err.(type) {
	case *PathError:
		{ //如果是自己定义的struct类型的错误
			fmt.Println("发生读错误", v)
		}
	default:
	}
}
