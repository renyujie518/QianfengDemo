package main

import (
	"fmt"
)

func main() {
	// 二维数组
	var value = [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	// 遍历二维数组的其他方法，使用 range
	// 其实，这里的 i, j 表示行游标和列游标
	// v2 就是具体的每一个元素
	// v  就是每一行的所有元素
	for i, v := range value {
		for j, v2 := range v {
			fmt.Printf("value[%v][%v]=%v \t ", i, j, v2)
		}
		fmt.Print(v)
		fmt.Println()
	}
}
