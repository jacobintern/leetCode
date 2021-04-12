---
title: 'Palindrome Number'
disqus: hackmd
---

## Beginners Guide

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".


Example 1:
---
```go=
Input: strs = ["flower","flow","flight"]
Output: "fl"
```

Example 2:
---
```go=
Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
```

Rules:
---
* 0 <= strs.length <= 200
* 0 <= strs[i].length <= 200
* strs[i] consists of only lower-case English letters.