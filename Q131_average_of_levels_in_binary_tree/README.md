# Average of Levels in Binary Tree(Easy)

## Beginners Guide

Given the `root` of a binary tree, return the average value of the nodes on each level in the form of an array. Answers within 10$^{-5}$ of the actual answer will be accepted.

### Example 1

```mermaid
graph TD
A((3))-->B((9))
A-->C((20))-->D((15))
C-->E((7))
```

>Input: root = [3,9,20,null,null,15,7]
Output: [3.00000,14.50000,11.00000]
Explanation: The average value of nodes on level 0 is 3, on level 1 is 14.5, and on level 2 is 11.
Hence return [3, 14.5, 11].

### Example 2

```mermaid
graph TD
A((3))-->B((9))-->C((15))
B-->D((7))
A-->E((20))
```

>Input: root = [3,9,20,15,7]
Output: [3.00000,14.50000,11.00000]

---

### Rules

* The number of nodes in the tree is in the range [1, 10$^4$].
* -2$^{31}$ <= Node.val <= 2$^{31}$ - 1
