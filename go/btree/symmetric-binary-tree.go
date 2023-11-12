package btree

import "fmt"

func isSymmetric(root *Node) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(left *Node, right *Node) bool {
	if left == nil && right == nil {
		return true
	}
	if left != nil && right != nil && left.Val == right.Val {
		return true && isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
	}
	return false
}

func RunIsSymmetric() {
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
				},
			},
			Right: &Node{
				Val: 2,
				Left: &Node{
					Val: 4,
				},
				Right: &Node{
					Val: 3,
				},
			},
		},
	}
	fmt.Println(isSymmetric(tree.Root))
}
