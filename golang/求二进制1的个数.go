// Copyright 2018 pfwu. All rights reserved.
//
// @Author: pfwu
// @Date: 2018-04-18 23:44

package main

import "fmt"

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