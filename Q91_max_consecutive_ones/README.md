---
title: 'Max Consecutive Ones(Easy)'
disqus: hackmd
---

## Beginners Guide

Given a binary array nums, return the maximum number of consecutive 1's in the array.

Example 1:
---
```go=
Input: nums = [1,1,0,1,1,1]
Output: 3
Explanation: The first two digits or the last three digits are consecutive 1s. The maximum number of consecutive 1s is 3.
```

Example 2:
---
```go=
Input: nums = [1,0,1,1,0,1]
Output: 2
```

Rules:
---
* $1 <= nums.length <= 10^5$
* nums[i] is either 0 or 1.