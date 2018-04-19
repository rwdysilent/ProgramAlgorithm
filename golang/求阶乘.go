// Copyright 2018 pfwu. All rights reserved.
//
// @Author: pfwu
// @Date: 2018-04-19 11:03

package main

import "fmt"

//num 的 n 次方
func power(num, n int) int {
	ans := 1
	for n > 0 {
		ans *= num
		n--
	}
	return ans
}

func main() {
	fmt.Println(power(3, 5))
}
