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

//定义一个快速排序函数
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

//定义一个快速排序
func QuickSort(arr []int, l, r int) {
	if l < r {
		i := l //左指针
		j := r //右指针
		key := arr[i]
		for i < j {
			for i < j && arr[j] >= key { //从右往左，找出第一个比基准小的索引
				j--
			}
			arr[i] = arr[j]
			for i < j && arr[i] <= key { //从左往右，找出第一个比基准大的索引
				i++
			}
			arr[j] = arr[i]
		}
		arr[i] = key
		QuickSort(arr, l, i-1)
		QuickSort(arr, i+1, r)
		fmt.Println(arr)
	}
}

//定义一个堆排序
func Heapsort() {

}

//定义一个插入排序
func insertSort(arr []int, rlen int) {
	var tem int
	for i := 1; i < rlen; i++ {
		j := i - 1
		tem = arr[i]
		for j >= 0 && tem < arr[j] {
			arr[j+1] = arr[j]
			j--
			fmt.Println("j--", j)
			fmt.Println(arr)
		}
		arr[j+1] = tem
		fmt.Println(arr)
	}
	fmt.Println(arr)
}

//选择排序
func selectionSort(arr []int, rlen int) {
	for i := 0; i < rlen-1; i++ { //有序区
		minIndex := i
		for j := i + 1; j < rlen; j++ { //无序区
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	fmt.Println(arr)
}
