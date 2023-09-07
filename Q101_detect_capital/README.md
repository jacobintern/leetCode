# Detect Capitals(Easy)

## Beginners Guide

We define the usage of capitals in a word to be right when one of the following cases holds:

* All letters in this word are capitals, like "USA".
* All letters in this word are not capitals, like "leetcode".
* Only the first letter in this word is capital, like "Google".
Given a string word, return true if the usage of capitals in it is right.



Example 1:
---
```go=
Input: word = "USA"
Output: true
```

Example 2:
---
```go=
Input: word = "FlaG"
Output: false
```

Example 3:
---
```go=
Input: word = "leetcode"
Output: true
```

Rules:
---
* 1 <= word.length <= 100
* word consists of lowercase and uppercase English letters.