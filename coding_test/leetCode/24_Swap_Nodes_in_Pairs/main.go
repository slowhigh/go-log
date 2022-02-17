package main

import (
	"fmt"
)

//Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)

//Example 1:
// 1 -> 2 -> 3 -> 4
//        ↓
// 2 -> 1 -> 4 -> 3
//
// Input: head = [1,2,3,4]
// Output: [2,1,4,3]

// Example 2:
// Input: head = []
// Output: []

// Example 3:
// Input: head = [1]
// Output: [1]

//Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	
	temp := head
	head = head.Next
	temp.Next = swapPairs(head.Next)
	head.Next = temp

	return head
}

func main() {
	node5 := ListNode { Val: 5, Next: nil}
	node4 := ListNode { Val: 4, Next: &node5}
	node3 := ListNode { Val: 3, Next: &node4}
	node2 := ListNode { Val: 2, Next: &node3}
	node1 := ListNode { Val: 1, Next: &node2}

	for node := &node1; ; node = node.Next {
		fmt.Println(node.Val)

		if node.Next == nil {
			break
		}
	}

	for node := swapPairs(&node1); ; node = node.Next {
		fmt.Println(node.Val)

		if node.Next == nil {
			break
		}
	}
}