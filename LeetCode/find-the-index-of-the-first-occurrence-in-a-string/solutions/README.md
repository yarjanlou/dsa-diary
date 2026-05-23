## 🧩 String Search Algorithms (`strStr`)

### 📖 Overview

This project implements the classic **substring search** problem — finding the first occurrence of a substring (`needle`) inside a larger string (`haystack`).

It contains both **naive** (**brute-force**) and optimized (**KMP**) approaches written in **JavaScript** and **Go**, along with benchmark results.

---

### 🚀 Features

- ✅ Simple and clear implementation
- ✅ Solutions in Go and JavaScript
- ✅ No third-party libraries
- 🧠 Naive and optimized implementations
- 📊 Benchmark and performance tracking supported

---

### 🧮 Algorithm 1: Naive Search

**Description:**
The naive search algorithm checks for the substring match by iterating through each position of the `haystack` and comparing it with the `needle`.

**Time Complexity:** `O(n * m)`
**Space Complexity:** `O(1)`

Where:

- `n` → length of `haystack`
- `m` → length of `needle`

---

### 🧮 Algorithm 2: Optimized Search (KMP)

**Description:**
The KMP algorithm improves efficiency by avoiding redundant comparisons.
It preprocesses the pattern (`needle`) to build an **LPS** (**longest prefix suffix**) table, then uses it to skip characters that have already been matched.

**Time Complexity:** `O(n + m)`
**Space Complexity:** `O(m)`

---

### 📊 Performance

This project includes benchmarks comparing execution time between the **naive** substring search approach and the **KMP** algorithm.

KMP demonstrates significantly better scalability, especially with large strings or repetitive patterns, because it avoids re-checking characters that have already been matched.

In summary:

- **Naive** → gets slower quickly as input grows
- **KMP** → remains fast and efficient at scale

You can visualize the performance results here:

![Performance Comparison](./performance/go-kmp.png)


---
