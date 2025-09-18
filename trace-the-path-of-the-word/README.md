# Trace the Path of the Word in Go

Given a grid of letters, check if a word can be traced by moving up, down, left, or right from one letter to the next.

Write a function that returns the path as a list of [row, col] positions, or "Not present" if the word is not found/can’t be created.

## 📝 Notes

- The target word will never be longer than the grid of letters.
- Target word and the letters grid will be in upper case.

## ✨ Examples

```go
trace_word_path("BISCUIT", [
  ["B", "I", "T", "R"],
  ["I", "U", "A", "S"],
  ["S", "C", "V", "W"],
  ["D", "O", "N", "E"]
])
output = [[0, 0], [1, 0], [2, 0], [2, 1], [1, 1], [0, 1], [0, 2]]

trace_word_path("HELPFUL", [
  ["L","I","T","R"],
  ["U","U","A","S"],
  ["L","U","P","O"],
  ["E","F","E","H"]
])
output = "Not present"

trace_word_path("UKULELE", [
  ["N", "H", "B", "W"],
  ["E", "X", "A", "D"],
  ["L", "A", "U", "U"],
  ["E", "L", "U", "K"]
])
output = [[2, 3], [3, 3], [3, 2], [3, 1], [3, 0], [2, 0], [1, 0]]

trace_word_path("SURVIVAL", [
  ["V", "L", "R", "L"],
  ["V", "A", "I", "V"],
  ["I", "O", "S", "C"],
  ["V", "R", "U", "F"]
])
output = [[2, 2], [3, 2], [3, 1], [3, 0], [2, 0], [1, 0], [1, 1], [0, 1]]
```
