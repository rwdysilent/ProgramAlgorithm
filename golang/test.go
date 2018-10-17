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

type ListNode1 struct {
	Node int
	Next *ListNode1
}

func reversList(node *ListNode1) *ListNode1 {
	if node == nil {
		return nil
	}

	var next *ListNode1
	q := node.Next
	for q != nil {
		node.Next = next
		next = node
		node = q
		q = node.Next
		if q == nil {
			node.Next = next
			return node
		}

	}
	return node
}

func Qsort(data []int) {
	if len(data) <= 1 {
		return
	}

	base := data[0]
	head, tail := 0, len(data)-1

	for i := 1; i <= tail; {
		if data[i] > base {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	Qsort(data[:head])
	Qsort(data[head+1:])
}

func main() {
	t := new(threadSafeSet)
	t.s = []string{"a", "b"}
	t.Iter()
}
