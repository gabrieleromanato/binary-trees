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
    public function minDepth(Node $root = null):  int {
        if (is_null($root)) {
            return 0;
        }

        if (is_null($root->left) && is_null($root->right)) {
            return 1;
        }

        if (!$root->left) {
            return $this->minDepth($root->right) + 1;
        }

        if (!$root->right) {
            return $this->minDepth($root->left) + 1;
        }

        return min($this->minDepth($root->left), $this->minDepth($root->right)) + 1;
    }
}

$root = new Node(1);
$root->left = new Node(9);
$root->right = new Node(20);
$root->right->left = new Node(15);
$root->right->right = new Node(7);

$solution = new Solution();
echo $solution->minDepth($root) . "\n"; // 2
