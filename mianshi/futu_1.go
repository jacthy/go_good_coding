package main

import (
	"errors"
	"fmt"
)

func main() {
	myMap := make(map[string]int)
	myMap["1"]=1
	myMap["2"]=2
	myMap["3"]=3
	//testZuHe()
}

func Append(s []int) {
	s = append(s,5)
}

func Add(s []int) {
	for _, val := range s {
		val += 5
	}
}

func Add2(s []int) {
	for i := 0; i <len(s); i++ {
		s[i] += 5
	}
}

func TiMu1() {
	arr := []int{1,2,3,4}
	fmt.Printf("%v\n",arr)
	Add(arr)
	fmt.Printf("%v\n",arr)
	Add2(arr)
	fmt.Printf("%v\n",arr)
}

type T interface {}

type S T

func TiMu2() {
	var (
		t T
		s S
		p *T
		i1 interface{}= t
		i2 interface{}= p
	)
	println(t == nil, p == nil, i1 == s)
	println(i1 == t, i1 == nil)
	println(i2 == p, i2 == nil)
}

func TiMu3()  {
	fmt.Println(foo())
}

func foo() (err error) {
	defer func() {
		fmt.Println(err)
		err = errors.New("a")
	}()
	defer func(e error) {
		fmt.Println(e)
		e = errors.New("b")
	}(err)
	return errors.New("c")
}

func testDefer() (t int){
	defer func() {
		println("1")
	}()
	defer func() {
		println("2")
	}()
	defer func(t int) {
		println(t)
	}(t)
	return 9
}

func countSum(arr []int) int {
	sum := 0
	for i := 0; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]
		if diff > 0 {
			sum += diff
		}
	}
	return sum
}
