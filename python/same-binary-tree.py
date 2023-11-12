class Node:
    def __init__(self, val):
        self.val = val
        self.left = None
        self.right = None

class Tree:
    def __init__(self, root):
        self.root = root


class Solution:
    def sameTree(self, root1, root2):
        if root1 is None and root2 is None:
            return True
        if root1 is not None and root2 is not None:
            return root1.val == root2.val and self.sameTree(root1.left, root2.left) and self.sameTree(root1.right, root2.right)
        return False
    

root1 = Node(1)
root1.right = Node(2)
root1.right.left = Node(3)

root2 = Node(1)
root2.right = Node(2)
root2.right.left = Node(3)

print(Solution().sameTree(root1, root2))