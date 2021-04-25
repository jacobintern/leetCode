---
title: 'Same Tree(easy)'
disqus: hackmd
---

## Beginners Guide

Given the roots of two binary trees p and q, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.


Example 1:
---
```go=
Input: p = [1,2,3], q = [1,2,3]
Output: true
```

Example 2:
---
```go=
Input: p = [1,2], q = [1,null,2]
Output: false
```

Example 3:
---
```go=
Input: p = [1,2,1], q = [1,1,2]
Output: false
```

Rules:
---
* The number of nodes in both trees is in the range [0, 100].
* $-10^4 <= Node.val <= 10^4$