class Node:
    def __init__(self, val):
        self.val = val
        self.left = None
        self.right = None

class Tree:
    def __init__(self, root):
        self.root = root

class Solution:
    def isBalanced(self, root):
        if root is None:
            return True
        return abs(self.maxDepth(root.left)-self.maxDepth(root.right)) < 2 and self.isBalanced(root.left) and self.isBalanced(root.right)

    def maxDepth(self, root):
        if root is None:
            return 0
        return max(self.maxDepth(root.left), self.maxDepth(root.right))+1
    
tree = Tree(Node(3))
tree.root.left = Node(9)
tree.root.right = Node(20)
tree.root.right.left = Node(15)
tree.root.right.right = Node(7)

print(Solution().isBalanced(tree.root))