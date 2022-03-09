package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	dirname := "/Users/renyujie/go/src/QianfengDemo"
	listfile(dirname, 0)

}
func listfile(dirname string, level int) {
	//level用来记录当前递归的层次，生成带有层次的空格
	s := "|-"
	for i := 0; i < level; i++ {
		s = "| " + s
	}
	fileinfo, err := ioutil.ReadDir(dirname)
	if err != nil {
		println(err)
		return
	}
	for _, info := range fileinfo {
		filename := dirname + "/" + info.Name()
		fmt.Printf("%s%s\n", s, filename)
		if info.IsDir() {
			//递归调用
			listfile(filename, level+1)
		}

	}

}
