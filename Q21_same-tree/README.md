# Same Tree(Easy)

## Beginners Guide

Given the roots of two binary trees `p` and `q`, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

### Example 1

```mermaid
graph TD

A((1)) --> B((2))
A --> C((3))

D((1)) --> E((2))
D --> F((3))

```

> Input: p = [1,2,3], q = [1,2,3]
Output: true

### Example 2

```mermaid
graph TD

A((1)) --> B((2))
A --> C((nil))

D((1)) --> E((nil))
D --> F((2))

```

> Input: p = [1,2], q = [1,null,2]
Output: false

### Example 3

```mermaid
graph TD

A((1)) --> B((2))
A --> C((1))

D((1)) --> E((1))
D --> F((2))

```

> Input: p = [1,2,1], q = [1,1,2]
Output: false

---

### Rules

* The number of nodes in both trees is in the range `[0, 100]`.
* -10$^4$ <= Node.val <= 10$^4$
