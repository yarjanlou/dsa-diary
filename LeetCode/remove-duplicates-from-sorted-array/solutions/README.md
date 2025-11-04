## ⚙️ How It Works

1. Initialize a pointer `j` at index `0` — this tracks the position of the last unique element.
2. Iterate through the array with index `i` starting from `1`.
3. Whenever `nums[i] != nums[j]`,

   * Move `nums[i]` to `nums[j+1]` (the next unique position).
   * Increment `j` by 1.
4. After finishing the loop, `j + 1` represents the count of unique elements.

---

## ⏱️ Time and Space Complexity

| Complexity Type | Explanation                                                                    |
| --------------- | ------------------------------------------------------------------------------ |
| **Time**        | `O(n)` — Each element is visited once.                                         |
| **Space**       | `O(1)` — The algorithm modifies the array in-place without using extra memory. |

---

