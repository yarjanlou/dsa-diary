# 🧩 Problem 66: Plus One — Solutions

## 🧠 Solution Overview

The JavaScript and Go solutions both scan the array from right to left and
propagate a carry only while they encounter `9`s.

There is no need to convert the digits into a numeric type. Avoiding that
conversion is important because the represented integer may be larger than the
language's safely supported integer range.

---

## ⚙️ How It Works

1. Start at the least significant digit, at the end of the array.
2. If the digit is smaller than `9`, increment it and return immediately.
3. Otherwise, replace the `9` with `0` and continue toward the beginning.
4. If the loop finishes, every digit was `9`, so return a result with a leading
   `1` followed by zeros.

### Walkthrough

For `digits = [1, 2, 9, 9]`:

```text
[1, 2, 9, 9]
          ↑  9 becomes 0; carry continues
[1, 2, 9, 0]
       ↑     9 becomes 0; carry continues
[1, 2, 0, 0]
    ↑        2 becomes 3; return
[1, 3, 0, 0]
```

---

## JavaScript

The JavaScript implementation updates the input array in place. When every
digit is `9`, it prepends `1` to the resulting zeros.

```javascript
var plusOne = function (digits) {
  for (let i = digits.length - 1; i >= 0; i--) {
    if (digits[i] + 1 != 10) {
      digits[i] = digits[i] + 1;
      return digits;
    }
    digits[i] = 0;
  }

  if (digits[0] == 0) {
    digits = [1, ...digits];
  }

  return digits;
};
```

---

## Go

The Go implementation also updates the slice in place when possible. If all
digits are `9`, it allocates a slice one element longer. Go initializes the
remaining elements to zero, so only the leading `1` must be assigned.

```go
func plusOne(digits []int) []int {
    for i := len(digits) - 1; i >= 0; i-- {
        if digits[i] < 9 {
            digits[i]++
            return digits
        }
        digits[i] = 0
    }
    res := make([]int, len(digits)+1)
    res[0] = 1
    return res
}
```

---

## ⏱️ Complexity

For both implementations:

- **Time:** O(n) in the worst case and O(1) in the best case
- **Space:** O(1) auxiliary space when no new leading digit is needed; O(n) for
  the returned array when the input consists entirely of `9`s

---

## ✅ Key Points

- The solutions correctly handle a single digit, trailing `9`s, and an input
  consisting entirely of `9`s.
- They avoid converting the digit array to an integer.
- They return as soon as the carry has been resolved.
