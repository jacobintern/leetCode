# Linked List Cycle(Easy)

## Beginners Guide

Given the `head` of a linked list, return the node where the cycle begins. If there is no cycle, return `null`.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the `next` pointer. Internally, `pos` is used to denote the index of the node that tail's `next` pointer is connected to **(0-indexed)**. It is `-1` if there is no cycle. **Note that** `pos` **is not passed as a parameter.**

**Do not modify** the linked list.

### Example 1

```mermaid
graph LR
A((3)) --> B((2)) --> C((0)) --> D((4))
D --> B
```

> Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).

### Example 2

```mermaid
graph LR
A((1)) --> B((2))
B --> A
```

> Input: head = [1,2], pos = 0
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.

### Example 3

```mermaid
graph LR
A((1))
```

> Input: head = [1], pos = -1
Output: false
Explanation: There is no cycle in the linked list.

---

### Rules

* The number of the nodes in the list is in the range [0, 10$^4$].
* -10$^5$ <= Node.val <= 10$^5$
* `pos` is `-1` or a **valid index** in the linked-list.
