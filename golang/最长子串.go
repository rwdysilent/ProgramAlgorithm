package main

import (
	"fmt"
	_"strings"
)

func lengthOfLongestSubstring(s string){
    max := 0
    hash_map := make(map[byte]int)
    // start unrepeated char
    start := 0
    for i := 0; i < len(s); i++ {
        if prev, ok := hash_map[s[i]]; ok {
            if start < prev + 1 {
                start = prev + 1
            }
        }
        if i-start+1 > max {
            max = i - start + 1
        }
        hash_map[s[i]] = i
    }
	fmt.Println(hash_map)
}

func main() {
	//s := "nfpdmpi"
	s := "abbacd"
	lengthOfLongestSubstring(s)
}