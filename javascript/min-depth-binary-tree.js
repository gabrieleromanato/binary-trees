'use strict';

const minDepth = (root) => {
    if (!root) return 0;
    if (!root.left) return minDepth(root.right) + 1;
    if (!root.right) return minDepth(root.left) + 1;
    if (minDepth(root.left) < minDepth(root.right)) return minDepth(root.left) + 1;
    return minDepth(root.right) + 1;
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

console.log(minDepth(root)); // 2