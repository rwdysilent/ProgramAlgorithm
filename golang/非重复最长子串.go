// Copyright 2018 pfwu. All rights reserved.
//
// @Author: pfwu
// @Date: 2018-04-29 18:52

//https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

//Given a string, find the length of the longest substring without repeating characters.
//
//Examples:
//
//Given "abcabcbb", the answer is "abc", which the length is 3.
//
//Given "bbbbb", the answer is "b", with the length of 1.
//
//Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a
//subsequence and not a substring.

//思路：
// 1. 构造一个map存储字符串最后一次出现的位置下标，key为字符，value为下标
// 2. 遍历字符串，如果在map中存在，说明有重复字符串，重新计算start起始位置。对于abba这样的字符，遍历到第二个b时，start更新为map中第一个b的下标
//	value+1，即2，继续遍历到最后一个a，start = map[a] + 1, 则为1，那么start位置更新错误，针对这样的情况，在更新start值时，需要一个判断，start
//	值小于map[key] + 1 再更新
// 3. 最大子串长度max为 遍历位置 - start + 1，更新max时和上一次max对比，小于不更新

package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
    max := 0
    hashMap := make(map[byte]int)

    // start unrepeated char
    start := 0
    for i := 0; i < len(s); i++ {
        if prev, ok := hashMap[s[i]]; ok {
        	//如果map中存在，则为重复字符串，如果子串起始值start小于重复字符串+1值，更新子串start值(避免abba这样的case)
            if start < prev + 1 {
                start = prev + 1
            }
        }
        //max值等于遍历下标值减去子串起始值+1
        if i-start+1 > max {
            max = i - start + 1
        }
        //更新map value为最后一次出现的字符下标值
        hashMap[s[i]] = i
    }
	return max
}

func main() {
	s := "abbacd"
	fmt.Println(lengthOfLongestSubstring(s))
}