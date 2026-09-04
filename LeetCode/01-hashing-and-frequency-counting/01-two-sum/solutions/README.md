# Solutions

## Approach

This is the canonical Complement Lookup problem. Instead of checking every pair with a nested loop, walk the array once and use a hash map to remember every value seen so far, keyed by its index. Before storing the current number, check whether its complement (`target - num`) is already in the map — if it is, the pair is found.

## How It Works

1. Create a hash map from value → index.
2. Iterate through `nums` once, tracking the current index `i` and value `num`.
3. On each step, compute the complement `target - num` and look it up in the map.
   - If found, return `[i, index]`, the indices of the two matching values.
   - If not found, store `num → i` in the map and continue.
4. Since the problem guarantees exactly one solution, the loop always returns before it ends.

## Complexity

- Time: `O(n)` — a single pass over the array with O(1) average map lookups/inserts.
- Space: `O(n)` — the map can hold up to n-1 entries before the match is found.

## Performance

![Go benchmark](./performance/main-go.jpg)
![JavaScript benchmark](./performance/main-js.jpg)
