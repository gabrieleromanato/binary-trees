class Node:
    def __init__(self, val):
        self.val = val
        self.left = None
        self.right = None

class Solution:
    def inorderTraversal(self, root):
        if root is None:
            return []
        stack = []
        res = []
        while root or stack:
            while root:
                stack.append(root)
                root = root.left
            if stack:
                node = stack.pop()
                res.append(node.val)
                root = node.right
        return res
    
root = Node(1)
root.right = Node(2)
root.right.left = Node(3)
print(Solution().inorderTraversal(root))