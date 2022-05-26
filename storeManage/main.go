package main

import (
	"log"
	"runtime"
	"time"
)

type People struct {
	Name  string
	Value string
}

func getMap() [2]string {
	pp := []*People{
		{Name: "A", Value: "a"},
		{Name: "B", Value: "B"},
	}
	m := [2]string{pp[0].Value, pp[1].Value}
	for i, people := range pp {
		m[i] = people.Value
	}
	return m
}

func main() {
	log.Println("start.")
	f()

	log.Println("force gc.")
	runtime.GC() // 调用强制gc函数

	log.Println("done.")
	time.Sleep(1 * time.Hour) // 保持程序不退出
}

func f() {
	container := make([]int, 8)
	log.Println("> loop.")
	// slice会动态扩容，用它来做堆内存的申请
	for i := 0; i < 32*1000*1000; i++ {
		container = append(container, i)
	}
	log.Println("< loop.")
	// container在f函数执行完毕后不再使用
}
