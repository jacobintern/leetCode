# Ransom Note(Easy)

## Beginners Guide

Given two strings `ransomNote` and `magazine`, return `true` if `ransomNote` can be constructed by using the letters from `magazine` and false otherwise.

Each letter in `magazine` can only be used once in `ransomNote`.

### Example 1

```go=
Input: ransomNote = "a", magazine = "b"
Output: false
```

### Example 2

```go=
Input: ransomNote = "aa", magazine = "ab"
Output: false
```

### Example 3

```go=
Input: ransomNote = "aa", magazine = "aab"
Output: true
```

---

### Rules

* $1 <= ransomNote.length, magazine.length <= 10^5$
* ransomNote and magazine consist of lowercase English letters.
