package swap_nodes_in_pairs_24

import "testing"

func Test_swapPairs(t *testing.T) {
	node4 := ListNode{Val: 4, Next: nil}
	node3 := ListNode{Val: 3, Next: &node4}
	node2 := ListNode{Val: 2, Next: &node3}
	node1 := ListNode{Val: 1, Next: &node2}

	output := []int{2, 1, 4, 3}
	for node, index := swapPairs(&node1), 0; ; node, index = node.Next, index+1 {
		if node.Val != output[index] {
			t.Error("Wrong output")
		}

		if node.Next == nil {
			break
		}
	}
}
