package main

/*
	The following program can find out the most optimal solution for the interval selection problem.
	Time complexity: O (n log n)
*/

import (
	"fmt"
	"sort"
)

func intervalSlection() {
	// takes the number of interval
	var n int
	fmt.Scan(&n)
	arr := make([][]int, n)
	// scans the range of n intervals
	for i := range arr {
		arr[i] = make([]int, 2)
		fmt.Scan(&arr[i][0], &arr[i][1])
	}

	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i][1] < arr[j][1]
	})

	stack := make([][]int, 0, n)
	top := 0
	stack = append(stack, arr[0])
	for i := 1; i < n; i++ {
		if stack[top][0] < arr[i][1] {
			stack = append(stack, arr[i])
			top++
		}
	}
	// Stack contains the intervals to select.
	fmt.Println(stack)
}
