# Symmetric Tree(Easy)

## Beginners Guide

Given the `root` of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

### Example 1

```mermaid
graph TD

A((1)) --> B((2)) --> C((3))
B --> D((4))
A --> E((2)) --> F((4))
E --> G((3))

```

> Input: root = [1,2,2,3,4,4,3]
Output: true

### Example 2

```mermaid
graph TD

A((1)) --> B((2)) --> C((nil))
B --> D((3))
A --> E((2)) --> F((nil))
E --> G((3))

```

> Input: root = [1,2,2,null,3,null,3]
Output: false

---

### Rules

* The number of nodes in the tree is in the range `[1, 1000]`.
* -100 <= Node.val <= 100
