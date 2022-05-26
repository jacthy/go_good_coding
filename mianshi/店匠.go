package main

import (
	"strconv"
	"strings"
)

// 重构下面代码
//func Handle() error {
//	if !Operation1(){
//		return OPERATION1FAILED
//	}
//	if !Operation2(){
//		return OPERATION2FAILED
//	}
//	if !Operation3(){
//		return OPERATION3FAILED
//	}
//	if  !Operation4() {
//		return OPERATION4FAILED
//	}
//	// do
//	return nil
//}

func getIntPair(arr []int, total int) [][2]int {
	if len(arr) < 2 {
		return nil
	}
	var resultPair [][2]int
	if arr[0] <= total { //arr[0]<=total才必要遍历
		for i := 1; i < len(arr); i++ {
			if arr[0]+arr[i] == total {
				resultPair = append(resultPair, [2]int{arr[0], arr[i]})
			}
		}
	}
	resultPair = append(resultPair, getIntPair(arr[1:], total)...)
	return resultPair
}

// IpToInt ip转数值
func IpToInt(ip string) uint32 {
	strArr := strings.Split(ip, ".")
	resultInt := uint32(0)
	for i := 0; i < len(strArr); i++ {
		valueInt, err := strconv.ParseInt(strArr[i], 10, 0)
		if err != nil {
			println(err.Error())
			return 0
		}
		resultInt = resultInt << 8
		resultInt = resultInt + uint32(valueInt)
	}
	return resultInt
}

// IntToIp 数值转ip
func IntToIp(ipInt uint32) string {
	str := ""
	for i := 3; i >= 0; i-- {
		val := ipInt & 255
		if str == "" {
			str = strconv.Itoa(int(val))
		}else {
			str = strconv.Itoa(int(val)) + "." + str
		}
		ipInt = ipInt >> 8
	}
	return str
}
