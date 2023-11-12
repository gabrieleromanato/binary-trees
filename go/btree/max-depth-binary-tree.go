package btree

import "fmt"

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func RunMaxDepth() {
	tree := Tree{
		Root: &Node{
			Val: 1,
			Left: &Node{
				Val: 2,
				Left: &Node{
					Val: 3,
				},
				Right: &Node{
					Val: 4,
					Left: &Node{
						Val: 5,
					},
				},
			},
			Right: &Node{
				Val: 2,
				Left: &Node{
					Val: 4,
					Right: &Node{
						Val: 5,
					},
				},
				Right: &Node{
					Val: 3,
				},
			},
		},
	}
	fmt.Println(maxDepth(tree.Root))
}
