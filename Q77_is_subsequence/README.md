---
title: 'Is Subquence(Easy)'
disqus: hackmd
---

## Beginners Guide

Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).

Example 1:
---
```go=
Input: s = "abc", t = "ahbgdc"
Output: true
```

Example 2:
---
```go=
Input: s = "axc", t = "ahbgdc"
Output: false
```

Rules:
---
* 0 <= s.length <= 100
* $0 <= t.length <= 10^4$
* s and t consist only of lowercase English letters.