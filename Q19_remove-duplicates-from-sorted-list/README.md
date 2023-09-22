# Remove Duplicates from Sorted List(Easy)

## Beginners Guide

Given the `head` of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list **sorted** as well.

### Example 1

```mermaid
graph LR

A((1)) --> B((1)) --> C((2))

D((1)) --> E((2))
```

> Input: head = [1,1,2]
Output: [1,2]

### Example 2

```mermaid
graph LR

A((1)) --> B((1)) --> C((2)) --> D((3)) --> E((3))

F((1)) --> G((2)) --> H((3))
```

> Input: head = [1,1,2,3,3]
Output: [1,2,3]

---

### Rules

* The number of nodes in the list is in the range `[0, 300]`.
* -100 <= Node.val <= 100
* The list is guaranteed to be **sorted** in ascending order.
