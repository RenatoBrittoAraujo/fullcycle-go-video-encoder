package main

import "fmt"

func main() {
	arr := []int{123, 12, 312, 321, 541, 24}
	MergeSort(&arr)
	fmt.Println(arr)
}
