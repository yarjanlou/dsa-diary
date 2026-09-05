# Solutions

## Approach

Reverse the digits of the number and compare the reversed value to the original. Negative numbers are rejected immediately, since the minus sign already makes them asymmetric.

## How It Works

1. If `x` is negative, return `false` immediately.
2. Store the original value of `x` for comparison later.
3. Initialize `reverseNumber` to `0`.
4. While `x > 0`:
   - Extract the last digit with `x % 10`.
   - Append it to `reverseNumber`: `reverseNumber = reverseNumber * 10 + (x % 10)`.
   - Remove the last digit from `x` via integer division.
5. Compare `reverseNumber` to the original value.

## Complexity

- Time: `O(log₁₀ n)` — proportional to the number of digits.
- Space: `O(1)`

## Performance

![Go benchmark](./performance/go-naive.png)
![JavaScript benchmark](./performance/js-naive.png)
