# Zigzag Conversion(Medium)

## Beginners Guide

The string `"PAYPALISHIRING`" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

```go
0   4   8     12
P   A   H     N
1 3 5 7 9  11 13
A P L S I  I  G
2   6   10
Y   I   R
```

And then read line by line: `"PAHNAPLSIIGYIR"`

Write the code that will take a string and make this conversion given a number of rows:

> string convert(string s, int numRows);

### Example 1

> Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Explanation:
P - A - H - N
A P L S I I G
Y - I - R

### Example 2

> Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:
P - - I - - - N
A - L S - - I G
Y A - H R
P - - I

### Example 3

> Input: s = "A", numRows = 1
Output: "A"

---

### Rules

* 1 <= s.length <= 1000
* `s` consists of English letters (lower-case and upper-case), `,` and `.`.
* 1 <= numRows <= 1000
