'use strict';

const sameBinaryTree = (p, q) => {
    if (!p && !q) {
        return true;
    }
    
    if (!p || !q) {
        return false;
    }
    
    if (p.val !== q.val) {
        return false;
    }
    
    return sameBinaryTree(p.left, q.left) && sameBinaryTree(p.right, q.right);
};

const p = {
    val: 1,
    left: {
        val: 2,
        left: null,
        right: null
    },
    right: {
        val: 3,
        left: null,
        right: null
    },
};

const q = {
    val: 1,
    left: {
        val: 2,
        left: null,
        right: null
    },
    right: {
        val: 3,
        left: null,
        right: null
    }
};

console.log(sameBinaryTree(p, q)); // true