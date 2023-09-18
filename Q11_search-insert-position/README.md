# Search Insert Position(Easy)

## Beginners Guide

Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

### Example 1

```go=
Input: nums = [1,3,5,6], target = 5
Output: 2
```

### Example 2

```go=
Input: nums = [1,3,5,6], target = 2
Output: 1
```

### Example 3

```go=
Input: nums = [1,3,5,6], target = 7
Output: 4
```

### Example 4

```go=
Input: nums = [1,3,5,6], target = 0
Output: 0
```

### Example 5

```go=
Input: nums = [1], target = 0
Output: 0
```

---

### Rules

* $1 <= nums.length <= 10^4$
* $-10^4 <= nums[i] <= 10^4$
* nums contains distinct values sorted in ascending order.
* $-10^4 <= target <= 10^4$
