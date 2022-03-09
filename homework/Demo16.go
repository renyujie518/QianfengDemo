package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//协程测试，模拟四个售票口同时卖20张票  但由于出现共享资源ticket，所以会出现ticket值忽变的情况，打印出负值
var ticket = 20

//编写一个售票的函数，需要判断是否还有票
var mutex sync.Mutex  //创建一把锁头
var wg sync.WaitGroup //设置同步等待锁

func saleticket(name string) {
	defer wg.Done() //每个go结束计数-1
	rand.Seed(time.Now().UnixNano())
	for { //需要一直判断
		mutex.Lock() //上锁
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)
			fmt.Println(name, "售出： ", ticket)
			ticket--
		} else {
			mutex.Unlock() //条件不满足也解锁
			fmt.Println(name, "已售光")
			break //结束掉这个for循环
		}
		mutex.Unlock() //解锁
	}
}
func main() {
	wg.Add(4) //买票前加入gorutine
	go saleticket("售票口1")
	go saleticket("售票口2")
	go saleticket("售票口3")
	go saleticket("售票口4")
	//为了保证主程序不先执行完，睡一会
	//time.Sleep(5*time.Second)

	wg.Wait() //所有的go没进行完，main阻塞
}
