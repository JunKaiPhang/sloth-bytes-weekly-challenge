# Evaluating Simple Algebra in Go

This program solves simple algebraic equations for **`x`**.  
It supports only **addition (`+`)** and **subtraction (`-`)**, with the guarantee that the input string always has spaces between numbers and symbols.

## ✨ Examples

```go
evalAlgebra("2 + x = 19")    // Output: 17
evalAlgebra("4 - x = 1")     // Output: 3
evalAlgebra("x + 10 = 53")   // Output: 43
evalAlgebra("-23 + x = -20") // Output: 3
evalAlgebra("10 + x = 5")    // Output: -5
evalAlgebra("-49 - x = -180")// Output: 131
evalAlgebra("x - 22 = -56")  // Output: -34
```

## 📌 Notes

1. There are spaces between every number and symbol in the string.

   - ✅ Example: `2 + x = 19`
   - ❌ Example: `2+x=19`

2. `x` may be a negative number.

## 🚀 How to Use

1. Run the program:

```bash
go run .
```

2. Enter the algebraic equation, then press Enter:

```shell
Enter equation: 2 + x = 19
Result: 17
```

Done ✅
