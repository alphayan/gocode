package main

import "fmt"

func main() {
	// BubbleSort(20)
	r := Rand(10)
	fmt.Println(r)
	//QuickSort(r, 0, 9)
	selectionSort(r, len(r))
	//fmt.Println(r)
	//HsetRedis()
	//HgetRedis()
	//defer RedisClose()
	//Hmset(FindMysql())
	//Hgetall()
	//RedisClose()
	//Test()
	//Goruntime1()
}
