# Repeated Substring Pattern(Easy)

## Beginners Guide

Given a string s, check if it can be constructed by taking a substring of it and appending multiple copies of the substring together.

### Example 1

```go=
Input: s = "abab"
Output: true
Explanation: It is the substring "ab" twice.
```

### Example 2

```go=
Input: s = "aba"
Output: false
```

### Example 3

```go=
Input: s = "abcabcabcabc"
Output: true
Explanation: It is the substring "abc" four times or the substring "abcabc" twice.
```

---

### Rules

* $1 <= s.length <= 10^4$
* s consists of lowercase English letters.
