# Solutions

## Approach

XOR every number in the array together. Since `a ^ a = 0` and `a ^ 0 = a`, every duplicated value cancels itself out, leaving only the number that appears once.

## How It Works

1. Initialize a running result to `0`.
2. XOR each element of the array into the result.
3. After the loop, the result is the single number — all duplicates have cancelled out.

## Complexity

- Time: `O(n)` — one pass over the array.
- Space: `O(1)` — no extra memory beyond the running result.

## Other Approaches

### Hash map (frequency counting)

Count the occurrences of each number in a hash map, then scan the map for the key with a count of `1`.

- Time: `O(n)`
- Space: `O(n)` — the map can hold up to `n` distinct values.

This is more intuitive but uses linear extra space; XOR achieves the same result with constant space, so it's the preferred approach here.

## Performance

![Go — naive (hash map)](./performance/go-naive.png)
![Go — XOR](./performance/go-xor.png)
