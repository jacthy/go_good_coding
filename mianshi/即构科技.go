package main

import "errors"

// 实现队列功能
type myQue struct {
	arr    []int
	pWrite int
	pPop   int
}

func newMyQue(len int) *myQue {
	arr := make([]int, len)
	return &myQue{arr: arr}
}

func (m *myQue) push(val int) error {
	willing := m.getWillingIndex(m.pWrite)
	if willing == m.pPop {
		return errors.New("overflow")
	}
	m.arr[willing] = val
	m.pWrite = willing
	return nil
}

func (m *myQue) pop() (int, error) {
	if m.pWrite == m.pPop {
		return 0, errors.New("there is no thing to pop")
	}
	result := m.arr[m.pPop]
	m.pPop = m.getWillingIndex(m.pPop)
	return result, nil
}

func (m *myQue) getWillingIndex(val int) int {
	if val == len(m.arr)-1 {
		return 0
	}
	return val + 1
}

// 数学计算题： RSA加密比ECC慢10倍，如果一台机器能处理6000个/秒，用RSA加密后能处理2000个/秒，那用ECC能处理多少个

// 算法题： a-z字符构成的长度为10的字符串，有26^10个组合。要求用0-26^10-1表示这些组合，写出func str2id()和func id2str()

func str2id(str string) {

}
