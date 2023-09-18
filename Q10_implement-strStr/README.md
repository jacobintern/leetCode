# Implement strStr()(easy)

## Beginners Guide

Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().

Example 1:

---

```go=
Input: haystack = "hello", needle = "ll"
Output: 2
```

Example 2:

---

```go=
Input: haystack = "aaaaa", needle = "bba"
Output: -1
```

Example 3:

---

```go=
Input: haystack = "", needle = ""
Output: 0
```

Rules:

---

* $0 <= haystack.length, needle.length <= 5 * 10^4$
* 0haystack and needle consist of only lower-case English characters.
