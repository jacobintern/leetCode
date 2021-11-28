---
title: 'Zigzag Conversion(Medium)'
disqus: hackmd
---

## Beginners Guide

The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

```
0   4   8     12
P   A   H     N
1 3 5 7 9  11 13
A P L S I  I  G
2   6   10
Y   I   R
```
And then read line by line: `PAHNAPLSIIGYIR`

Write the code that will take a string and make this conversion given a number of rows:

```
string convert(string s, int numRows);
```

Example 1:
---
```go=
Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Explanation:
0   4   8     12
P   A   H     N
1 3 5 7 9  11 13
A P L S I  I  G
2   6   10
Y   I   R
```

Example 2:
---
```go=
Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:
0     6       12
P     I       N
1   5 7    11 13
A   L S    I  G
2 4   8 10
Y A   H R
3     9
P     I
```

Example 3:
---
```go=
Input: s = "A", numRows = 1
Output: "A"
```

Rules:
---
* $1 <= s.length <= 1000$
* `s` consists of English letters (lower-case and upper-case), `,` and `.`.
* 1 <= numRows <= 1000