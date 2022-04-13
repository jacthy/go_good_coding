package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main1() {
	ch :=make(chan int)
	wg := sync.WaitGroup{}
	// 启动1个生产者
	wg.Add(1)
	go func() {
		wg.Done()
		producer(ch)
	}()
	// 启动3个consumer
	for i := 0; i < 3; i++ {
		go consumer(ch)
	}

	for  {
	}
}

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		println("producer produce val: ", i)
		ch <- i
		time.Sleep(300*time.Millisecond)
	}
}

func consumer(ch chan int) {
	for val := range ch {
		println("consumer get val:", val)
	}
}




// 使用goroutine和channel实现一个计算int64随机数各位和的程序

type job struct {
	val int64
}

type result struct {
	*job
	sum int64
}

// producer 生产者：负责产生数据并将其放入缓冲区jobChan, jobChan是一个只读通道
func producer2(jobChan chan<- *job, wg *sync.WaitGroup) {
	defer wg.Done()
	// 1. 开启一个goroutine循环生成int64随机数，发送到jobChan
	for{
		x := rand.Int63()
		// 生产者产生“货物”
		newJob := &job{
			val: x,
		}
		jobChan <- newJob
		// goroutine运行太快，在此减慢一下速度
		time.Sleep(time.Microsecond * 100)
	}
}

// consumer 消费者：负责从jobChan缓冲区中读出数据，对数据处理后写入resultChan缓冲区
func consumer2(jobChan <-chan *job, resultChan chan<- *result, wg *sync.WaitGroup) {
	// 2. 从jobChan中取出随机数计算个位数的和，将结果发送到resultChan
	defer wg.Done()
	for{
		// 读取原数据
		job := <- jobChan
		var sum int64 = 0
		n := job.val
		// 计算各位和
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &result{
			job: job,
			sum: sum,
		}
		resultChan <- newResult
	}
}

func main2() {
	// 声明缓冲区
	var jobChan = make(chan *job, 100)
	var resultChan = make(chan *result, 100)
	// 线程同步
	var wg sync.WaitGroup
	wg.Add(1)
	go producer2(jobChan, &wg)
	// 开启24个goroutine
	wg.Add(24)
	for i := 0; i < 24; i++{
		go consumer2(jobChan, resultChan, &wg)
	}
	// 3. 主goroutine从resultChan取出结果并打印在console
	i := 0
	for result := range resultChan{
		i++
		fmt.Printf("value: %d sum: %d count: %d\n", result.job.val, result.sum, i)
	}
	wg.Wait()
}


// -----------------------------
func main4() {
	var wg sync.WaitGroup
	ch := make(chan int, 3)
	workerNum := 3
	for i := 1; i <= workerNum; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			consumer3(id, ch)
		}(i)
	}
	go producer3(ch)
	wg.Wait()
	fmt.Printf("work done \n")
}

func consumer3(id int, ch chan int) {
	for data := range ch {
		fmt.Printf("workerId:%d consumer data:%d \n", id, data)
	}
	fmt.Printf("workerId:%d consumer finish \n", id)
}

func producer3(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
	fmt.Printf("producer finish \n")
}