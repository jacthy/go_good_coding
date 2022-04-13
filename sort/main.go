package main

import "fmt"

func main()  {
	arr := []int{2,3,1,6,2,7,8,7,4,5}
	fmt.Printf("%v" , mergeSort(arr))
}

func mergeSort(arr []int) []int {
	l := len(arr)
	if l == 1 {
		return arr
	}
	mid := l/2
	return merge(mergeSort(arr[:mid]),mergeSort(arr[mid:]))
}

func merge(left, right []int) []int {
	l,r := 0,0
	result := make([]int,0,l+r)
	for l < len(left) || r < len(right) {
		if l == len(left) {
			result= append(result, right[r:]...)
			break
		}else if r == len(right) {
			result= append(result, left[l:]...)
			break
		}else if left[l] < right[r] {
			result= append(result, left[l])
			l++
		}else {
			result= append(result, right[r])
			r++
		}
	}
	return result
}