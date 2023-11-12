<?php

class Node {
    public $val;
    public $left;
    public $right;

    public function __construct($val) {
        $this->val = $val;
        $this->left = null;
        $this->right = null;
    }
}

class Tree {
    public $root;
}

class Solution {
    public static function isSameTree($p, $q) {
        if(is_null($p) && is_null($q)) {
            return true;
        }
        if(is_null($p) || is_null($q)) {
            return false;
        }
        if($p->val != $q->val) {
            return false;
        }
        return self::isSameTree($p->left, $q->left) && self::isSameTree($p->right, $q->right);
    }
}

$tree = new Tree();
$tree->root = new Node(1);
$tree->root->left = new Node(2);
$tree->root->right = new Node(3);

$tree2 = new Tree();
$tree2->root = new Node(1);
$tree2->root->left = new Node(2);
$tree2->root->right = new Node(3);

var_dump(Solution::isSameTree($tree->root, $tree2->root));