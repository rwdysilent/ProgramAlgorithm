// Copyright 2018 pfwu. All rights reserved.
//
// @Author: pfwu
// @Date: 2018-03-24 22:41

package main

import "fmt"

func f(x int) func(y int) int{
	x++
	var sum = func(y int) int {
		return x + y
	}
	return sum
}

func main(){
	a := f(1)

	fmt.Println(a(2))
	fmt.Println(a(1))
}