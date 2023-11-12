package btree

import "fmt"

func inOrderTraversal(node *Node) []int {
	var result []int
	var stack []*Node
	if node == nil {
		return result
	}
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		if len(stack) > 0 {
			node = stack[len(stack)-1]
			result = append(result, node.Val)
			stack = stack[:len(stack)-1]
			node = node.Right
		}

	}
	return result
}

func RunInOrderTraversal() {
	tree := Tree{
		Root: &Node{
			Val: 1,
			Left: &Node{
				Val: 2,
			},
			Right: &Node{
				Val: 3,
				Left: &Node{
					Val: 4,
				},
			},
		},
	}
	fmt.Println(inOrderTraversal(tree.Root))
}
