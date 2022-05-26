package main

import "fmt"

type tree struct {
	left  *tree
	right *tree
	val   int
}

func getTreeBySlice(n int, arr []int) *tree {
	if len(arr) <= 1 {
		return nil
	}
	fmt.Printf("%v,%v\n", n, arr[0])
	getTreeBySlice(n+1, arr[get2Min(n):])
	getTreeBySlice(n+1, arr[get2Min(n)+1:])
	return nil
}

func get2Min(n int) int {
	min := 1
	for i := 0; i < n-1; i++ {
		min = min * 2
	}
	return min
}

type treeSlice []int

func (t treeSlice) node(index int) (val, left, right int) {
	val = t[index]
	left = index*2 + 1
	right = index*2 + 2
	return
}

func forwardPrint(index int, t treeSlice) {
	if index >= len(t) {
		return
	}
	val, left, right := t.node(index)
	forwardPrint(left,t)
	println(val)
	forwardPrint(right,t)
}
