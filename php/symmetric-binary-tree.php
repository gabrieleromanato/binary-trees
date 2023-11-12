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
    public static function isSymmetric(Node $root): bool {
        if(is_null($root)) {
            return true;
        }
        return self::isMirror($root->left, $root->right);
    }

    public static function isMirror(Node $left = null, Node $right = null): bool {
        if(is_null($left) && is_null($right)) {
            return true;
        }
        if(is_null($left) || is_null($right)) {
            return false;
        }
        return ($left->val == $right->val) && self::isMirror($left->left, $right->right) && self::isMirror($left->right, $right->left);
    }
}

$root = new Node(1);
$root->left = new Node(2);
$root->right = new Node(2);
$root->left->left = new Node(3);
$root->left->right = new Node(4);
$root->right->left = new Node(4);
$root->right->right = new Node(3);

var_dump(Solution::isSymmetric($root));