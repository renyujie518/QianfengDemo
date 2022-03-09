package main

import "fmt"

func feibonaqi(n int) int {
	if n < 2 {
		return n
	}
	return feibonaqi(n-2) + feibonaqi(n-1)
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(feibonaqi(i))
	}

}
