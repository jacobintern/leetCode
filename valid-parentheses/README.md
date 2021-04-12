---
title: 'Palindrome Number(easy)'
disqus: hackmd
---

## Beginners Guide

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.


Example 1:
---
```go=
Input: s = "()"
Output: true
```

Example 2:
---
```go=
Input: s = "()[]{}"
Output: true
```

Example 3:
---
```go=
Input: s = "(]"
Output: false
```

Example 4:
---
```go=
Input: s = "([)]"
Output: false
```

Example 5:
---
```go=
Input: s = "{[]}"
Output: true
```

Rules:
---
* $1 <= s.length <= 10^4$
* s consists of parentheses only '()[]{}'.