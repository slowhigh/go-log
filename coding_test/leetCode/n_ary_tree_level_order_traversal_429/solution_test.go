package n_ary_tree_level_order_traversal_429

import (
	"fmt"
	"testing"
)

// Input: root = [1,null,3,2,4,null,5,6]
// Output: [[1],[3,2,4],[5,6]]

type testCase struct {
	input *Node
	output [][]int
}

func Test_levelOrder(t *testing.T) {
	e1n6 := Node {
		Val: 6,
		Children: []*Node {},
	}
	e1n5 := Node {
		Val: 5,
		Children: []*Node {},
	}
	e1n4 := Node {
		Val: 4,
		Children: []*Node {},
	}
	e1n2 := Node {
		Val: 2,
		Children: []*Node {},
	}
	e1n3 := Node {
		Val: 3,
		Children: []*Node {
			&e1n5,
			&e1n6,
		},
	}
	e1n1 := Node {
		Val: 1,
		Children: []*Node {
			&e1n3,
			&e1n2,
			&e1n4,
		},
	}

	testCaseArr := []testCase {
		{
			input: &e1n1,
			output: [][]int {
				{1},
				{3,2,4},
				{5,6},
			},
			
		},
		{
			input: nil,
			output: [][]int {},
		},
	}

	for _, tc := range testCaseArr {
		if fmt.Sprint(levelOrder(tc.input)) != fmt.Sprint(tc.output) {
			t.Error("wrong Error")
		} 
	}
}