'use strict';

const binaryTreeInorderTraversal = (root) => {
    const result = [];
    const stack = [];
    let current = root;
    
    while (current || stack.length) {
        if (current) {
        stack.push(current);
        current = current.left;
        } else {
        current = stack.pop();
        result.push(current.val);
        current = current.right;
        }
    }
    
    return result;
};

const root = {
    val: 1,
    left: null,
    right: {
        val: 2,
        left: {
        val: 3,
        left: null,
        right: null,
        },
        right: null,
    },
};

console.log(binaryTreeInorderTraversal(root)); // [1, 3, 2]