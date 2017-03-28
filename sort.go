package main

import (
	"fmt"
	"math/rand"
	"time"
)

//定义一个随机函数用来生成数组
var array []int

func Rand(n int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		array = append(array, r.Intn(100))
	}
	return array
}

//定义一个冒泡排序函数
func BubbleSort(n int) {
	arr := Rand(n)
	fmt.Println("冒泡排序前：", arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("冒泡排序后", arr)
}

func Quick2Sort(values []int) {
	if len(values) <= 1 {
		return
	}
	mid, i := values[0], 1
	head, tail := 0, len(values)-1
	for head < tail {
		//fmt.Println(values)
		if values[i] > mid {
			values[i], values[tail] = values[tail], values[i]
			tail--
		} else {
			values[i], values[head] = values[head], values[i]
			head++
			i++
		}
	}
	values[head] = mid
	Quick2Sort(values[:head])
	Quick2Sort(values[head+1:])
}
func QuickSort(arr []int, l, r int) {
	if l < r {
		i := l //左指针
		j := r //右指针
		key := arr[i]
		for i < j {
			for i < j && arr[j] > key {
				j--
			}
			arr[i] = arr[j]
			for i < j && arr[i] < key {
				i++
			}
			arr[j] = arr[i]
		}
		array[i] = key
		QuickSort(array, l, i-1)
		QuickSort(array, i+1, r)

	}
}

func Heapsort() {

}
