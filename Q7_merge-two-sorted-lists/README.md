# Merge Two Sorted Lists(Easy)

## Beginners Guide

You are given the heads of two sorted linked lists `list1` and `list2`.

Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

### Example 1

```mermaid
graph LR
A((1)) --> B((2)) --> C((4))
D((1)) --> E((3)) --> F((4))
G((1)) --> H((1)) --> I((2)) --> J((3)) --> K((4)) --> L((4))

style A fill:red,color:white
style B fill:red,color:white
style C fill:red,color:white
style D fill:purple,color:white
style E fill:purple,color:white
style F fill:purple,color:white
style G fill:purple,color:white
style H fill:red,color:white
style I fill:red,color:white
style J fill:purple,color:white
style K fill:red,color:white
style L fill:purple,color:white
```

> Input: l1 = [1,2,4], l2 = [1,3,4]
Output: [1,1,2,3,4,4]

### Example 2

> Input: l1 = [], l2 = []
Output: []

### Example 3

> Input: l1 = [], l2 = [0]
Output: [0]

---

### Rules

* The number of nodes in both lists is in the range [0, 50].
* -100 <= Node.val <= 100
* Both `list1` and `list2` are sorted in non-decreasing order.
