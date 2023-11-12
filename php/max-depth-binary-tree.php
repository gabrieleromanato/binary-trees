<?php

class Node {
    public $data;
    public $left;
    public $right;

    public function __construct($data) {
        $this->data = $data;
        $this->left = null;
        $this->right = null;
    }
}

class Tree {
    public $root;
}

class Solution {
    public static function maxDepth($root) {
        if (is_null($root)) {
            return 0;
        }

        $leftDepth = self::maxDepth($root->left);
        $rightDepth = self::maxDepth($root->right);

        return max($leftDepth, $rightDepth) + 1;
    }
}

$tree = new Tree();
$tree->root = new Node(3);
$tree->root->left = new Node(9);
$tree->root->right = new Node(20);
$tree->root->right->left = new Node(15);
$tree->root->right->right = new Node(7);

echo Solution::maxDepth($tree->root); // 3