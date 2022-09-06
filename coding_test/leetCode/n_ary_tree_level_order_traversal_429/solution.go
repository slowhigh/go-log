package n_ary_tree_level_order_traversal_429

// Definition for a Node.
type Node struct {
    Val int
    Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int {}
	}

	temp := []*Node { root }
	result := make([][]int, 0)

	for ; len(temp) != 0; {
		childrenArr := make([]*Node, 0)
		valArr := make([]int, 0)
		for _, v := range temp {
			if v == nil {
				continue
			}
			childrenArr = append(childrenArr, v.Children...)
			valArr = append(valArr, v.Val)
		}

		temp = childrenArr
		result = append(result, valArr)
	}

	return result
}