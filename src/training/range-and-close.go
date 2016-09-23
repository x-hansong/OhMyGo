package main

import (
	"fmt"
)

func fibonacci(n int, c chan int)  {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	//循环 `for i := range c` 会不断从 channel 接收值，直到它被关闭。
	for i := range c {
		fmt.Println(i)
	}
}
