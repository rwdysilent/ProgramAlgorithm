// Copyright 2018 pfwu. All rights reserved.
//
// @Author: pfwu
// @Date: 2018-05-09 13:47

// 给定一个数n，打印n之前的菲波那切数列

package main

import "fmt"

func fibonacci(c, quit chan int, n int) {
	x, y := 0, 1
	for {
		select {
		case <-c:
			x, y = y, x+y
			if x > n {
				return
			}
			fmt.Println(x)
		case <-quit:
			return
		}
	}
}

func main() {
	c, q := make(chan int), make(chan int)
	n := 10
	go func() {
		for i := 0; i < n; i++ {
			c <- 1
		}
		q <- 1
	}()
	fibonacci(c, q, n)
}
