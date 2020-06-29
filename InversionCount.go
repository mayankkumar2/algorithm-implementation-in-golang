package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

/*
	The following program is implementation of Inversion Counter using Merge Sort
	Complexity : O(n log n)
*/

var count = 0

func MergeSort(arr *[]int, initFlag int, finalFlag int) {
	length := finalFlag - initFlag + 1
	if length < 2 {
		return
	} else {

		MergeSort(arr, initFlag, initFlag+((length-1)/2))
		MergeSort(arr, initFlag+((length-1)/2)+1, finalFlag)
		Merge(arr, initFlag, finalFlag, initFlag+((length-1)/2))
	}
}
func Merge(arr *[]int, beg int, end int, mid int) {
	subs := make([]int, end-beg+1)
	i := 0
	p1 := beg
	p2 := mid + 1
	for i < (end - beg + 1) {
		if (*arr)[p1] < (*arr)[p2] {
			subs[i] = (*arr)[p1]
			i++
			p1++
		} else if (*arr)[p1] > (*arr)[p2] {
			subs[i] = (*arr)[p2]
			count += mid - p1 + 1
			i++
			p2++
		}
		if p1 > mid || p2 > end {
			break
		}
	}
	for p1 <= mid {
		subs[i] = (*arr)[p1]
		i++
		p1++
	}

	for p2 <= end {
		subs[i] = (*arr)[p2]
		i++
		p2++
	}
	for i, v := range subs {
		(*arr)[beg+i] = v
	}
}

func InversionCounter() int {
	body, ok := ioutil.ReadFile("./dataset")
	if ok != nil {
		fmt.Println(ok.Error())
		return 0
	}
	arrayOfStrings := strings.Split(string(body), "\n")
	arr := make([]int, len(arrayOfStrings))
	for i := range arrayOfStrings {
		arr[i], _ = strconv.Atoi(arrayOfStrings[i])
	}
	MergeSort(&arr, 0, len(arr)-1)
	return count
}
