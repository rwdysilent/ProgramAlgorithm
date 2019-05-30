// Copyright 2019 wupengfei. All rights reserved.
//
// @Author: wupengfei
// @Date: 2019-05-29 20:50

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseListNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var pre *ListNode
	cur := head

	for cur != nil {
		next := cur.Next

		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

func createListNode(s []int, n int) *ListNode {
	if n == 0 {
		return nil
	}

	var head = ListNode{
		Val: s[0],
	}

	curNode := &head
	for i := 1; i < n; i++ {
		next := new(ListNode)
		next.Val = s[i]
		curNode.Next = next
		curNode = curNode.Next
	}

	return &head
}

func printListNode(head *ListNode) {
	for head != nil {
		fmt.Printf("val: %d\n", head.Val)
		head = head.Next
	}
	fmt.Println("End......")
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	n := createListNode(s, 5)
	printListNode(n)

	r := reverseListNode(n)
	printListNode(r)
}
