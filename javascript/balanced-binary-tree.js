'use strict';

const maxDepth = (root) => {
    if (!root) return 0;
    return Math.max(maxDepth(root.left), maxDepth(root.right)) + 1;
};

const isBalanced = (root) => {
    if (!root) return true;
    return Math.abs(maxDepth(root.left) - maxDepth(root.right)) <= 1 
    && isBalanced(root.left) && isBalanced(root.right);
};

const root = {
    val: 3,
    left: {
        val: 9,
        left: null,
        right: null,
    },
    right: {
        val: 20,
        left: {
            val: 15,
            left: null,
            right: null,
        },
        right: {
            val: 7,
            left: null,
            right: null
        }
    }
};

console.log(isBalanced(root)); // true