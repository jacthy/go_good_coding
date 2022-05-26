package main

import "fmt"

// ONES程诚面试题 一个链表穿插的拆成两个链表

type node struct {
	data int
	next *node
}

type ctx struct {
	first  *node
	second *node
}

func split(head *node) ctx {
	ctxCon := ctx{
		first:  nil,
		second: nil,
	}
	// 代码在这里开始写
	if head == nil {
		return ctxCon
	}
	ctxCon.first = head
	if head.next == nil {
		return ctxCon
	}
	ctxCon.second = head.next
	for head != nil {
		tmp := head.next
		if head.next != nil {
			head.next = head.next.next
		}
		head = tmp
	}
	return ctxCon
}

func testSplit() {
	tmp := new(node)
	head := tmp
	// 生成链
	for i := 1; i < 9; i++ {
		newNode := new(node)
		newNode.data = i
		tmp.next = newNode
		tmp = newNode
	}
	ctx := split(head)
	printNode(ctx.first)
	printNode(ctx.second)
}

func printNode(head *node) {
	for head != nil {
		fmt.Printf("%d,", head.data)
		head = head.next
	}
	println()
}
