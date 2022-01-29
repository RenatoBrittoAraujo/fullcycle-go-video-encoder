package main

import (
	"fmt"

	"github.com/renatobrittoaraujo/fullcycle-go-video-encoder/sorting"
)

func main() {
	arr := []int{123, 12, 312, 321, 541, 24}
	sorting.MergeSort(&arr)
	fmt.Println(arr)
}
