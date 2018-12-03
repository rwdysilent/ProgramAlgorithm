// Copyright 2018 pfwu. All rights reserved.
//
// @Author: pfwu
// @Date: 2018-04-18 23:44

//任意给定一个32位无符号整数n，求n的二进制表示中1的个数，比如n = 5（0101）时，返回2，n = 15（1111）时，返回4
//对于二进制数来说，除一个2，就少一位，可以判断这个少的位来确定"1"的个数。

package main

import (
	"fmt"
)

func count(n byte) {
	IntNum := 0
	for n > 0 {
		if n%2 == 1 {
			IntNum++
		}
		n = n / 2
	}
	fmt.Print(IntNum)
}

func main() {
	count(15)
}