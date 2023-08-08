---
title: 'Quest Name(Level)'
disqus: hackmd
---

## Beginners Guide

Given an array of strings words, return the words that can be typed using letters of the alphabet on only one row of American keyboard like the image below.

In the American keyboard:

* the first row consists of the characters "qwertyuiop",
* the second row consists of the characters "asdfghjkl", and
* the third row consists of the characters "zxcvbnm".

Example 1:
---
```go=
Input: words = ["Hello","Alaska","Dad","Peace"]
Output: ["Alaska","Dad"]
```

Example 2:
---
```go=
Input: words = ["omk"]
Output: []
```

Example 3:
---
```go=
Input: words = ["adsdf","sfd"]
Output: ["adsdf","sfd"]
```

Rules:
---
* 1 <= words.length <= 20
* 1 <= words[i].length <= 100
* words[i] consists of English letters (both lowercase and uppercase). 