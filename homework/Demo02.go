package main

import "fmt"

//打印1-100内，能够被3整除，但是不能被5整除的数字，统计个数同时每行打印5个
func main() {
	count := 0
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 != 0 {
			fmt.Print(i, "\t")
			count++
			if count%5 == 0 {
				fmt.Println()
			}
		}
	}
	fmt.Println()
	fmt.Print("个数总数为", count)

}
