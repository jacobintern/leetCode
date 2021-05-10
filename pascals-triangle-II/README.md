---
title: 'Pascal's Triangle II(easy)'
disqus: hackmd
---

## Beginners Guide

Given an integer rowIndex, return the rowIndexth (0-indexed) row of the Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:

        | 1 |
      | 1 | 1 |
    | 1 | 2 | 1 |
  | 1 | 3 | 3 | 1 | 
| 1 | 4 | 6 | 4 | 1 |


Example 1:
---
```go=
Input: rowIndex = 3
Output: [1,3,3,1]
```

Example 2:
---
```go=
Input: rowIndex = 0
Output: [1]
```

Example 3:
---
```go=
Input: rowIndex = 1
Output: [1,1]
```

Rules:
---
* 1 <= rowIndex <= 33