package main

import (
	"fmt"
)

func Goruntime() {
	slic := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	ch := make([]chan int, 10) //定义channel切片
	for i := 0; i < 10; i++ {
		ch[i] = make(chan int)
		go func(c []chan int, s []int, j int) {
			fmt.Println("j", s[j])
			c[j] <- j
		}(ch, slic, i)
	}

	for k, _ := range ch {
		fmt.Println("ch", <-ch[k])
	}
}
func Goruntime1() {
	slic := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	m := 0
	ch := make(chan int, 10) //定义缓冲channel
	for i := 0; i < 10; i++ {
		go func(c chan int, s []int, j int) {
			fmt.Println("j", s[j])
			c <- i
		}(ch, slic, i)

	}
	//带缓冲的channel需要关闭，要不然会死锁
	for i := range ch {
		m++
		fmt.Println("ch", i)
		if m >= 10 {
			close(ch)
		}
	}

}
