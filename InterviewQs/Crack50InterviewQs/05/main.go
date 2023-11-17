package main

import "fmt"

func Merge(left, rght []int) []int {
	merged := make([]int, 0, len(left)+len(rght))
	for len(left) > 0 || len(rght) > 0 {
		if len(left) == 0 {
			return append(merged, rght...)
		} else if len(rght) == 0 {
			return append(merged, left...)
		} else if left[0] < rght[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, rght[0])
			rght = rght[1:]
		}
	}

	return merged
}

func MergeSort_SEQ(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	midd := len(data) / 2
	left := MergeSort(data[:midd])
	rght := MergeSort(data[midd:])
	return Merge(left, rght)
}

func MergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	done := make(chan bool)
	midd := len(data) / 2
	var left []int
	go func() {
		left = MergeSort(data[:midd])
		done <- true
	}()
	rght := MergeSort(data[midd:])
	<-done
	return Merge(left, rght)
}

func main() {
	println("beg")
	data := []int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8}
	//fmt.Printf("%v\n%v\n", data, MergeSort_SEQ(data))
	fmt.Printf("%v\n%v\n", data, MergeSort(data))
	println("end")
}
