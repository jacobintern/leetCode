# Binary tree tilt(Easy)

## Beginners Guide

Given the `root` of a binary tree, return the sum of every tree node's **tilt**.

The **tilt** of a tree node is the **absolute difference** between the sum of all left subtree node **values** and all right subtree node **values**. If a node does not have a left child, then the sum of the left subtree node **values** is treated as `0`. The rule is similar if the node does not have a right child.

### Example 1

```mermaid
graph TD
A((1)) --> B((2))
A --> C((3))

D((1)) --> E((0))
D --> F((0))
```

> Input: root = [1,2,3]
Output: 1
Explanation:
Tilt of node 2 : |0-0| = 0 (no children)
Tilt of node 3 : |0-0| = 0 (no children)
Tilt of node 1 : |2-3| = 1 (left subtree is just left child, so sum is 2; right subtree is just right child, so sum is 3)
Sum of every tilt : 0 + 0 + 1 = 1

### Example 2

```mermaid
graph TD
A((4)) --> B((2)) --> C((3))
B --> D((5))
A --> E((9)) --> F((nil))
E --> G((7))

H((6)) --> I((2)) --> J((0))
I --> K((0))
H --> L((7)) --> M((nil))
L --> N((0))
```

> Input: root = [4,2,9,3,5,null,7]
Output: 15
Explanation:
Tilt of node 3 : |0-0| = 0 (no children)
Tilt of node 5 : |0-0| = 0 (no children)
Tilt of node 7 : |0-0| = 0 (no children)
Tilt of node 2 : |3-5| = 2 (left subtree is just left child, so sum is 3; right subtree is just right child, so sum is 5)
Tilt of node 9 : |0-7| = 7 (no left child, so sum is 0; right subtree is just right child, so sum is 7)
Tilt of node 4 : |(3+5+2)-(9+7)| = |10-16| = 6 (left subtree values are 3, 5, and 2, which sums to 10; right subtree values are 9 and 7, which sums to 16)
Sum of every tilt : 0 + 0 + 0 + 2 + 7 + 6 = 15

### Example 3

```mermaid
graph TD
A((21)) --> B((7)) --> C((1)) --> D((3))
C --> E((3))
B --> F((1))
A --> G((14)) --> H((2))
G --> Z((2))

J((3)) --> K((6)) --> L((0)) --> M((0))
L --> N((0))
K --> O((0))
J --> P((0)) --> Q((0))
P --> R((0))
```

> Input: root = [21,7,14,1,1,2,2,3,3]
Output: 9

---

### Rules

* The number of nodes in the tree is in the range [0, 10$^4$].
* -1000 <= Node.val <= 1000
