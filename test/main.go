package main

import (
	"fmt"
	"time"
)

type L struct {
	v1 int
	v2 string
}

func main() {
	fmt.Printf("%v\n", 1 << 10)

}

func change2(m *map[string]int) {

	//fmt.Printf("\nmm:%v\n",m)
}
func change(arr []int)  {
	arr[2]=4
}

func testGo() {
	go func() {
		println("start")
		go func() {
			for true {
				println("still working")
				time.Sleep(500 * time.Millisecond)
			}
		}()
		println("end")
		return
	}()
}