# Unique Binary Search Trees(Medium)

## Beginners Guide

Given an integer `n`, return the number of structurally unique **BST**'s (binary search trees) which has exactly `n` nodes of unique values from `1` to `n`.

### Example 1

```mermaid
graph TD
A((1))-->B((nil))
A-->C((3))-->D((2))
C-->E((nil))

F((1))-->G((nil))
F-->H((2))-->I((nil))
H-->J((3))

K((2))-->L((1))
K-->M((3))

N((3))-->O((2))-->P((1))
N-->Q((nil))
O-->R((nil))

S((3))-->T((1))
S-->U((nil))
T-->V((nil))
T-->W((2))
```

>Input: n = 3
Output: 5

### Example 2

>Input: n = 1
Output: 1

---

### Rules

* `1 <= n <= 19`
