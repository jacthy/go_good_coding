package main

import "go/types"

type student struct {

}

type linkedMap struct {
	types.Map
}

type teacher struct {
	student
}

func (s student) sayA() {
	println("sA")
}

func (s student) sayB() {
	println("sB")
}

func (t teacher) sayA() {
	println("tA")
}

func testZuHe()  {
	new(teacher).sayA()
	new(teacher).sayB()
}
