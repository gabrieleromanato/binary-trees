package btree

import "fmt"

func isSameTree(p *Node, q *Node) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil && p.Val == q.Val {
		return true && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return false
}

func RunIsSameTree() {
	tre1 := Tree{
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
	tre2 := Tree{
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

	fmt.Println(isSameTree(tre1.Root, tre2.Root))
}
