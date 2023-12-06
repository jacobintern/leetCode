# Swap Nodes in Pairs(Medium)

## Beginners Guide

Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)

### Example 1

```mermaid
graph LR
A((1))-->B((2))-->C((3))-->D((4))

to

E((2))-->F((1))-->G((4))-->H((3))
```

>Input: head = [1,2,3,4]
Output: [2,1,4,3]

### Example 2

>Input: head = []
Output: []

### Example 3

>Input: head = [1]
Output: [1]

---

### Rules

* The number of nodes in the list is in the range `[0, 100]`.
* `0 <= Node.val <= 100`
