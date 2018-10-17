// Copyright 2018 pfwu. All rights reserved.
//
// @Author: pfwu
// @Date: 2018-04-28 23:33

//n 个面包，两种吃法，一次一个或者一次两个，求出所有可能的序列

//TODO: 未完善

package main

import (
	"fmt"
	)

var EndAns [][]string

func eatBread(n int) {
	var ans []string
	if n < 2 {
		ans = append(ans, "1")
		fmt.Println(ans)
		return
	}
	for i := 0; i < n/2; i++ {
		ans = append(ans, "2")
	}
	if n%2 == 0 {
		fmt.Println(ans)
		return
	}
	ans = append(ans, "1")
	fmt.Println(ans)
}

func arrangement(ans []string) {
	tmp := []string{"1", "1"}
	for i := range ans {
		if ans[i] == "2" {
			var out []string
			out = append(out, ans[:i]...)
			out = append(out, tmp...)
			out = append(out, ans[i+1:]...)
			fmt.Println(out)
			arrangement(out)
		}
	}
}

func fab(n int) {
	var s string = ""
	myFab(n, s)
}

func myFab(n int, s string) {
	if n == 0 {
		fmt.Println(s)
		return
	}
	if n > 0 {
		myFab(n-1, s+"1")
	}
	if n > 1 {
		myFab(n-2, s+"2")
	}
}

func main() {
	//eatBread(4)
	//arrangement([]string{"2", "2"})
	//fmt.Println("End: ", EndAns)
	//
	//fmt.Println(strings.Count("fivefivesasd", "ive"))
	//fmt.Println(strings.HasPrefix("-&-weixin", "&"))
	//fmt.Println(strings.IndexAny("-&weixin", "#"))

	myFab(10, "")
}
