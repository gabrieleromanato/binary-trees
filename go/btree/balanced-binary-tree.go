package btree

import "math"

func height(root *Node) float64 {
	if root == nil {
		return 0
	}
	return 1 + math.Max(height(root.Left), height(root.Right))
}

func isBalanced(root *Node) bool {
	if root == nil {
		return true
	}
	return math.Abs(height(root.Left)-height(root.Right)) < 2 && isBalanced(root.Left) && isBalanced(root.Right)
}

func RunIsBalanced() {
	tree := Tree{
		Root: &Node{
			Val: 3,
			Left: &Node{
				Val: 9,
			},
			Right: &Node{
				Val: 20,
				Left: &Node{
					Val: 15,
				},
				Right: &Node{
					Val: 7,
				},
			},
		},
	}
	println(isBalanced(tree.Root))
}
