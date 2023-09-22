# Convert Sorted Arrat to Binary Search Tree(Level)

## Beginners Guide

Given an integer array (nums) wheretje elements are sorted in **ascending order**, convert it to a height-balanced binary search tree.

### Example 1

```mermaid
graph TD

A((0)) --> B(( -3 )) --> C(( -10 ))
B --> D((nil))
A --> E((9)) --> F((5))
E --> G((nil))

H((0)) --> I(( -10 )) --> J((nil))
I --> K(( -3 ))
H --> L((5)) --> M((nil))
L --> N((9))

```

> Input: nums = [-10,-3,0,5,9]
Output: [0,-3,9,-10,null,5]
Explanation: [0,-10,5,null,-3,null,9] is also accpeted.

### Example 2

```mermaid
graph TD

A((3)) --> B((1))
A --> C((nil))

D((2)) --> E((nil))
D --> F((3))

```

> Input: nums = [1,3]
Output: [3,1]
Explanation: [1,null,3] or [3,1] are both height-balanced BSTs.

---

### Rules

* 1 <= nums.length <= 10$^4$
* -10^4 <= nums[i] <= 10$^4$
* `nums` is sorted in a **strictly increasing** order.
