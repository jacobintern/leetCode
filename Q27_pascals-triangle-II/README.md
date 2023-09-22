# Pascal's Triangle II(Easy)

## Beginners Guide

Given an integer `rowIndex`, return the rowIndex$^{th}$ (**0-indexed**) row of the **Pascal's triangle**.

In **Pascal's triangle**, each number is the sum of the two numbers directly above it as shown:

```go
        | 1 |
      | 1 | 1 |
    | 1 | 2 | 1 |
  | 1 | 3 | 3 | 1 | 
| 1 | 4 | 6 | 4 | 1 |
```

### Example 1

> Input: rowIndex = 3
Output: [1,3,3,1]

### Example 2

> Input: rowIndex = 0
Output: [1]

### Example 3

> Input: rowIndex = 1
Output: [1,1]

---

### Rules

* 1 <= rowIndex <= 33
