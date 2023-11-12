package btree

import (
	"fmt"
	"math"
)

func minDepth(root *Node) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}

	return int(math.Min(float64(minDepth(root.Left)), float64(minDepth(root.Right))) + 1)
}

func RunMinDepth() {
	tree := Tree{
		Root: &Node{
			Val: 1,
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
	fmt.Println(minDepth(tree.Root))
}
