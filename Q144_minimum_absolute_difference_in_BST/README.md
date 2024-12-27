# Minimum Absolute Difference in BST(Easy)

## Beginners Guide

Given the `root` of a Binary Search Tree (BST), return *the minimum absolute difference between the values of any two different nodes in the tree*.

### Example 1
```mermaid
graph TD
A((4))-->B((2))-->C((3))
B-->D((3))
A-->E((6))
```

>Input: root = [4,2,6,1,3]
Output: 1

### Example 2
```mermaid
graph TD
A((4))-->B((2))-->C((3))
B-->D((3))
A-->E((6))
```

>Input: root = [1,0,48,null,null,12,49]
Output: 1

### Example 3
```mermaid
graph TD
A((236))-->B((104))-->C((nil))
B-->D((227))
A-->E((701))-->F((nil))
E-->G((911))
```

>Input: root = [236,104,701,null,227,null,911]
Output: 9

---

### Rules

* The number of nodes in the tree is in the range [2, 10$^4$].
* 0 <= Node.val <= 10$^5$
