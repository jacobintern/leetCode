---
title: 'Ugly Number(Easy)'
disqus: hackmd
---

## Beginners Guide

An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.

Given an integer n, return true if n is an ugly number.

Example 1:
---
```go=
Input: n = 6
Output: true
Explanation: 6 = 2 × 3
```

Example 2:
---
```go=
Input: n = 1
Output: true
Explanation: 1 has no prime factors, therefore all of its prime factors are limited to 2, 3, and 5.
```

Example 3:
---
```go=
Input: n = 14
Output: false
Explanation: 14 is not ugly since it includes the prime factor 7.
```

Rules:
---
* $-2^31 <= n <= 2^31 - 1$