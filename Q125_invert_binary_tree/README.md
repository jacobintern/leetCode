# Invert Binary Tree(Easy)

## Beginners Guide

Given the `root` of a binary tree, invert the tree, and return its root.

### Example 1

```mermaid
graph TD
A((4))-->B((2))-->C((1))
B-->D((3))
A-->E((7))-->F((6))
E-->G((9))

to

H((4))-->I((7))-->J((9))
I-->K((6))
H-->L((2))-->M((3))
L-->N((1))
```

>Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]

### Example 2

```mermaid
graph TD
A((2))-->B((1))
A-->C((3))

to

D((2))-->E((3))
D-->F((1))
```

>Input: root = [2,1,3]
Output: [2,3,1]

### Example 3

>Input: root = []
Output: []

---

### Rules

* The number of nodes in the tree is in the range `[0, 100]`.
* -100 <= Node.val <= 100
