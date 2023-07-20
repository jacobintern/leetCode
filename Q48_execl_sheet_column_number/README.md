---
title: 'Excel Sheet Column Number(Easy)'
disqus: hackmd
---

## Beginners Guide

Given a string columnTitle that represents the column title as appears in an Excel sheet, return its corresponding column number.

For example:
```
A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28 
...
```

Example 1:
---
```go=
Input: columnTitle = "A"
Output: 1
```

Example 2:
---
```go=
Input: columnTitle = "AB"
Output: 28
```

Example 3:
---
```go=
Input: columnTitle = "ZY"
Output: 701
```

Rules:
---
* 1 <= columnTitle.length <= 7
* columnTitle consists only of uppercase English letters.
* columnTitle is in the range ["A", "FXSHRXW"].