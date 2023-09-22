# Path Sum(Easy)

## Beginners Guide

Given the `root` of a binary tree and an integer `targetSum`, return `true` if the tree has a **root-to-leaf** path such that adding up all the values along the path equals `targetSum`.

A leaf is a node with no children.

### Example 1

```mermaid
graph TD

A((5)) --> B((4)) --> C((11)) --> D((7))
C --> E((2))
A --> F((8)) --> G((8)) --> H((13))
G --> I((4)) --> J((nil))
I --> K((1))

style A fill:#7fc3ff
style B fill:#7fc3ff
style C fill:#7fc3ff
style E fill:#7fc3ff
```

> Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
Output: true

### Example 2

```mermaid
graph TD

A((1)) --> B((2))
A --> C((3))

```

> Input: root = [1,2,3], targetSum = 5
Output: false

### Example 3

> Input: root = [1,2], targetSum = 0
Output: false

---

### Rules

* The number of nodes in the tree is in the range `[0, 5000]`.
* -1000 <= Node.val <= 1000
* -1000 <= targetSum <= 1000
