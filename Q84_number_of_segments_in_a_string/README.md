---
title: 'Number of Segments in a String(Easy)'
disqus: hackmd
---

## Beginners Guide

Given a string s, return the number of segments in the string.

A segment is defined to be a contiguous sequence of non-space characters.

Example 1:
---
```go=
Input: s = "Hello, my name is John"
Output: 5
Explanation: The five segments are ["Hello,", "my", "name", "is", "John"]
```

Example 2:
---
```go=
Input: s = "Hello"
Output: 1
```

Rules:
---
* 0 <= s.length <= 300
* s consists of lowercase and uppercase English letters, digits, or one of the following characters "!@#$%^&*()_+-=',.:".
* The only space character in s is ' '.