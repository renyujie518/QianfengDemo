package main

import (
	"fmt"
	"sort"
)

func main() {
	map1 := make(map[int]string)
	map1[1] = "Demo1"
	map1[2] = "Demo2"
	map1[3] = "Demo3"
	map1[6] = "Demo6"
	fmt.Println(map1)
	for k, v := range map1 {
		fmt.Println(k, "-->", v) //map是无序存储的
	}
	fmt.Println("-------")
	//要想获得顺序的map，要先把key排个序,在切片里排序
	keys := make([]int, 0, len(map1))
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println("经过排序后的map为 :")
	for _, value := range keys {
		fmt.Printf("%d-->%s\n", value, map1[value])

	}
	//map结合slice
	map01 := map[string]string{"name": "renyujie", "age": "22", "sex": "male"}
	map02 := map[string]string{"name": "yuwenge", "age:": "23", "sex": "female"}
	map_slice := make([]map[string]string, 0, len(map01))
	map_slice = append(map_slice, map01)
	map_slice = append(map_slice, map02)
	for i, v := range map_slice {
		fmt.Println("信息汇总")
		fmt.Printf("第%d人， 姓名是 %s，年龄是 %s，性别是 %s\n", i,
			v["name"], v["age"], v["sex"])

	}

}
