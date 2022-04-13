package main

import (
	"fmt"
	"time"
)

//func main() {
//	test2()
//}

func test1() {
	// 死锁1
	ch := make(chan int)
	ch <- 100
	num := <-ch
	fmt.Println("num=", num)
}

func test2() {

	ch := make(chan int)
	//ch <- 100 //此处死锁 优于go程之前使用通道
	go func() {
		num := <-ch
		fmt.Println("num=", num)
	}()
	ch <- 100  // 此处不死锁
	time.Sleep(time.Second * 3)
	fmt.Println("finish")

}
