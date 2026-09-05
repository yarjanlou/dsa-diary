# Solutions

## Approach

Process the digits from right to left, the way manual addition works: increment the last digit, and only carry into earlier digits when the current one was a `9`. Neither implementation converts the digit array into a numeric type, since the represented integer may exceed the language's safely representable integer range.

## How It Works

1. Start at the least significant digit, at the end of the array.
2. If the digit is smaller than `9`, increment it and return immediately.
3. Otherwise, replace the `9` with `0` and continue toward the beginning.
4. If the loop finishes, every digit was `9`, so return a new array with a leading `1` followed by zeros.

This naturally handles a single-digit input, trailing `9`s, and an input made entirely of `9`s.

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

## Complexity

- Time: `O(n)` worst case (all digits are `9`), `O(1)` best case.
- Space: `O(1)` auxiliary; `O(n)` for the returned array only when a new leading digit is needed.

## Performance

![Go benchmark](./performance/main-go.jpg)
![JavaScript benchmark](./performance/main-js.png)
