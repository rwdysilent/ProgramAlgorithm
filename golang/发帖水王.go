// Copyright 2018 pfwu. All rights reserved.
//
// @Author: pfwu
// @Date: 2018-04-19 21:22

package main

import "fmt"

var (
	ids = []int{1, 2, 1, 1, 1, 2, 2, 2, 3, 3, 3, 13, 3, 14, 15}
)

//Tango是微软亚洲研究院的一个试验项目。研究院的员工和实习生们都很喜欢在Tango上面交流灌水。传说，Tango有一大“水王”，
//他不但喜欢发贴，还会回复其他ID发的每个帖子。坊间风闻该“水王”发帖数目超过了帖子总数的一半。如果你有一个当 前论坛上
//所有帖子（包括回帖）的列表，其中帖子作者的ID也在表中，你能快速找出这个传说中的Tango水王吗？
func findOneId() int {
	ID := []int{1, 2, 1, 1, 5, 1, 1, 1, 9}
	candidate := 0
	times := 0
	for i := 1; i < len(ID); i++ {
		if times == 0 {
			candidate = ID[i]
			times = 1
		} else {
			if candidate == ID[i] {
				times++
			} else {
				times--
			}
		}
	}
	return candidate
}

//随着Tango的发展，管理员发现，“超级水王”没有了。统计结果表明，有3个发帖很多的ID，
//他们的发帖数目都超过了帖子总数目N的1/4。你能从发帖ID列表中快速找出他们的ID吗？
func findThreeIds([]int) []int {
	times := []int{0, 0, 0}
	candidate := []int{0, 0, 0}
	for i := 0; i < len(ids); i++ {
		if ids[i] == candidate[0] {
			times[0]++
		} else if ids[i] == candidate[1] {
			times[1]++
		} else if ids[i] == candidate[2] {
			times[2]++
		} else if times[0] == 0 {
			times[0] = 1
			candidate[0] = ids[i]
		} else if times[1] == 0 {
			times[1] = 1
			candidate[1] = ids[i]
		} else if times[2] == 0 {
			times[2] = 1
			candidate[2] = ids[i]
		} else {
			times[0]--
			times[1]--
			times[2]--
		}
	}
	fmt.Println(candidate)
	return candidate
}

func main() {
	fmt.Println(findOneId())
	findThreeIds(ids)
}
