// Copyright 2018 pfwu. All rights reserved.
//
// @Author: pfwu
// @Date: 2018-04-28 23:33

//n 个面包，两种吃法，一次一个或者一次两个，求出所有可能的序列
//动态规划：https://zhuanlan.zhihu.com/p/31628866

//TODO: 未完善

package main

import "fmt"

func FindN(n int) []string {
	if n == 0 {
		return []string{}
	}
	if n == 1 {
		return []string{"1"}
	}
	if n == 2 {
		return []string{"11", "2"}
	}

	var tmp []string

	var Get func(n int) []string
	Get = func(n int) []string {
		tmp = append([]string{"1"}, Get(n-1)...)
		tmp = append([]string{"2"}, Get(n-2)...)
		tmp = append([]string{"11"}, Get(n-2)...)
		return tmp
	}

	return tmp
}

func main() {
	t := FindN(2)
	fmt.Println("=========")
	fmt.Printf("%+v", t)
}
