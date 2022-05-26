package main

import "fmt"

func main() {
	arr := []int{1, 5, 3, 8, 9, 8, 6, 0, 4, 2}
	forwardPrint(0,treeSlice(arr))
}



// 发现个有趣的现像
func intesting() {
	arr := []int{1, 5, 3, 8, 9, 8, 6, 0, 4, 2}
	arr1 := arr[2:4:5]
	arr1=append(arr1, 10)//这里会时arr[4]变成10
	arr1[0]=0
	arr1[1] =0
	arr2 :=arr[10:] // arr[len:]不越界，只是返回空数组，arr[len+1:]才越界
	arr2 =arr[11:]
	fmt.Printf("%v\n",arr)
	fmt.Printf("%v\n",arr1)
	fmt.Printf("%v\n",arr2)
}

func mergeSort(arr []int) []int {
	l := len(arr)
	if l == 1 {
		return arr
	}
	mid := l / 2
	return merge(mergeSort(arr[:mid]), mergeSort(arr[mid:]))
}

func merge(left, right []int) []int {
	l, r := 0, 0
	result := make([]int, 0, l+r)
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}




func guibin(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	min := len(arr) / 2
	left := guibin(arr[0:min])
	right := guibin(arr[min:])
	return mergeArr(left, right)
}

func mergeArr(left, right []int) []int {
	l, r := len(left), len(right)
	lIndex, rIndex := 0, 0
	var tmpArr []int
	for lIndex < l && rIndex < r {
		if left[lIndex] < right[rIndex] {
			tmpArr = append(tmpArr, left[lIndex])
			lIndex++
		} else {
			tmpArr = append(tmpArr, right[rIndex])
			rIndex++
		}
	}
	tmpArr = append(tmpArr, left[lIndex:]...)
	tmpArr = append(tmpArr, right[rIndex:]...)
	return tmpArr
}

func insertSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if arr[j] > arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}
}

func shellSort(arr []int) {
	for gap := len(arr) / 2; gap > 0; gap = gap / 2 {
		for j := 0; j < len(arr)-gap; j = j + gap {
			for i := j + gap; i > 0; i = i - gap {
				if arr[i] > arr[i-gap] {
					arr[i], arr[i-gap] = arr[i-gap], arr[i]
				}
			}
		}
	}
}

func quickSort(arr []int) {
	if len(arr)<=1 {
		return
	}else if len(arr) == 2 {
		if arr[0] < arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
		}
		return
	}
	left :=  1
	right := len(arr) - 1
	for left <= right {
		if left == right {
			arr[0], arr[left] = arr[left], arr[0]
			break
		}
		if arr[right] > arr[0] {
			right--
			continue
		}
		if arr[left] < arr[0] {
			left++
			continue
		}
		arr[left], arr[right] = arr[right], arr[left]
	}
	quickSort(arr[:left])
	quickSort(arr[right+1:])
}
