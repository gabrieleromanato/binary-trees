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

class Solution {
    public $result = [];

    public function inorderTraversal($root) {
        $stack = [];
        $current = $root;
        while(!is_null($current) || count($stack) > 0) {
            if(!is_null($current)) {
                array_push($stack, $current);
                $current = $current->left;
            } else {
                $current = array_pop($stack);
                array_push($this->result, $current->val);
                $current = $current->right;
            }
        }
        return $this->result;
    }
}

$root = new Node(1);
$root->right = new Node(2);
$root->right->left = new Node(3);

$solution = new Solution();
print_r($solution->inorderTraversal($root));