package main

/*
	The following program is a bruteforce algorithm for testing and solving the 2 robot problem.
	Time complexity: O (2^N)
*/

import (
	"fmt"
	"math"
)

var n, m int
var arr [][]int

func calculate(bot1Pos int, bot2Pos int, distance int, c int) int {
	var b1distance, b2distance int
	if c >= n {
		return distance
	}
	if bot1Pos == -1 {
		b1distance = int(math.Abs(float64(arr[c][1]) - float64(arr[c][0])))
	} else {
		b1distance = int(math.Abs(float64(bot1Pos)-float64(arr[c][0])) + math.Abs(float64(arr[c][1])-float64(arr[c][0])))
	}

	if bot2Pos == -1 {
		b2distance = int(math.Abs(float64(arr[c][1]) - float64(arr[c][0])))
	} else {
		b2distance = int(math.Abs(float64(bot2Pos)-float64(arr[c][0])) + math.Abs(float64(arr[c][1])-float64(arr[c][0])))
	}

	ans := int(math.Min(float64(calculate(arr[c][1], bot2Pos, distance+b1distance, c+1)), float64(calculate(bot1Pos, arr[c][1], distance+b2distance, c+1))))
	return ans
}

func solve() int {
	fmt.Println(calculate(-1, -1, 0, 0))
	return -1
}

// TwoRobotProblemBruteforce is a one of the undesired solution as it takes 2^N time
// a better solution is a DP approach a top down - Memoization, or a or a bottom-up approach
func TwoRobotProblemBruteforce() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&m, &n)
		arr = make([][]int, n)
		for f := range arr {
			arr[f] = make([]int, 2)
			fmt.Scan(&arr[f][0], &arr[f][1])
		}
		solve()
	}
}
