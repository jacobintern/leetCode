# Binary Tree Preorder Traversal(Easy)

## Beginners Guide

Given the root of a binary tree, return the preorder traversal of its nodes' values.

### Example 1

```mermaid
graph TD

A((1))-->B((nil))
A-->C((2))-->D((3))
C-->E((nil))
```

>Input: root = [1,null,2,3]
Output: [1,2,3]

### Example 2

>Input: root = []
Output: []

### Example 3

>Input: root = [1]
Output: [1]

---

### Rules

* The number of nodes in the tree is in the range `[0, 100]`.
* -100 <= Node.val <= 100
