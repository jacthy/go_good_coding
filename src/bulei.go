package main

import "math/rand"

func buLei(n, m, k int) [][]int {
	arrList := make([][]int, n)
	for i := 0; i < len(arrList); i++ {
		for j := 0; j < m; j++ {
			arrList[i] = append(arrList[i], 0)
		}
	}
	bego := make([]int, 0, k)
	for len(bego) <= k {
		n := rand.Intn(n * m)
		if !check(bego, n) {
			bego = append(bego, n)
		}
	}
	for _, lei := range bego {
		row := lei / n
		col := lei % n
		arrList[row][col] = -1
		println("lei:", row, col)
		for _, position := range getShowCountPosition(lei, n, m) {
			println("leizhouwei:", position.row, position.col)
			arrList[position.row][position.col]++
		}
	}
	return arrList
}

type position struct {
	row int
	col int
}

func getShowCountPosition(n, row, col int) []position {
	i := n / row
	j := n % row
	rowUp, rowDown, colLeft, colRight := 0, 0, 0, 0
	if i == 0 {
		rowUp = i
		rowDown = i + 1
	} else if i == row-1 {
		rowUp = i - 1
		rowDown = i
	} else {
		rowUp = i - 1
		rowDown = i + 1
	}
	if j == 0 {
		colLeft = j
		colRight = j + 1
	} else if j == col-1 {
		colLeft = j - 1
		colRight = j
	} else {
		colLeft = j - 1
		colRight = j + 1
	}
	return getPositions(rowUp, rowDown, colLeft, colRight, i, j)
}

func getPositions(rowUp, rowDown, colLeft, colRight, i, j int) []position {
	var arr []position
	for x := rowUp; x <= rowDown; x++ {
		for y := colLeft; y <= colRight; y++ {
			if x == i && y == j {
				continue
			}
			arr = append(arr, position{row: x, col: y})
		}
	}
	return arr
}

func check(arr []int, num int) bool {
	for _, value := range arr {
		if value == num {
			return true
		}
	}
	return false
}
