package swap_nodes_in_pairs_24

type ListNode struct {
	Val  int
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
