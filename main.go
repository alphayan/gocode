package main

import (
	"fmt"
)

func main() {
	//BubbleSort(20)
	r := Rand(10)
	fmt.Println(r)
	QuickSort(r, 0, 9)
	fmt.Println(r)
}
