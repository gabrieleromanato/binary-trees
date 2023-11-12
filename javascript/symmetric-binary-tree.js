'use strict';

const isSymmetric = (root) => {
    const isMirror = (p, q) => {
        if (!p && !q) {
            return true;
        }
        
        if (!p || !q) {
            return false;
        }
        
        return p.val === q.val && isMirror(p.left, q.right) && isMirror(p.right, q.left);
    };
    
    return isMirror(root, root);
};

const root = {
    val: 1,
    left: {
        val: 2,
        left: {
        val: 3,
        left: null,
        right: null,
        },
        right: null,
    },
    right: {
        val: 2,
        left: null,
        right: {
        val: 3,
        left: null,
        right: null,
        },
    },
};

console.log(isSymmetric(root)); // true