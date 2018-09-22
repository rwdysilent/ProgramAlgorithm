// Copyright 2018 pfwu. All rights reserved.
//
// @Author: pfwu
// @Date: 2018-09-19 18:41

package main

import (
	"sync"
	"fmt"
)

type threadSafeSet struct {
	s []string
	sync.RWMutex
}

var wg sync.WaitGroup

func (set *threadSafeSet) Iter() {
	ch := make(chan interface{})
	//go func() {
	//	set.RLock()
	//	for elem := range set.s {
	//		ch <- elem
	//	}
	//	close(ch)
	//	set.RUnlock()
	//}()

	for elem := range set.s {
		wg.Add(1)
		go func(elem int) {
			defer wg.Done()
			ch <- elem
		}(elem)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for range ch {
		fmt.Println(<-ch)
	}
	//return ch
}

func main() {
	t := new(threadSafeSet)
	t.s = []string{"a", "b"}
	t.Iter()
}
