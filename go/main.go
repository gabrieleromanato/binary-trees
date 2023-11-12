package main

import (
	"fmt"
	"goinfoteam/btree"
)

func main() {
	fmt.Println("Binary Tree Inorder Traversal")
	btree.RunInOrderTraversal()
	fmt.Println("Same Binary Tree")
	btree.RunIsSameTree()
	fmt.Println("Symmetric Binary Tree")
	btree.RunIsSymmetric()
	fmt.Println("Max Depth Binary Tree")
	btree.RunMaxDepth()
	fmt.Println("Balanced Binary Tree")
	btree.RunIsBalanced()
	fmt.Println("Min Depth Binary Tree")
	btree.RunMinDepth()
}
