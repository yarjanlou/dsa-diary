## 🧩 String Search Algorithms (`strStr`)

### 📖 Overview

This project implements the classic **substring search** problem — finding the first occurrence of a substring (`needle`) inside a larger string (`haystack`).

Currently, it includes a **naive (brute-force)** implementation written in Go.
Future updates will add the **KMP algorithm** and benchmark comparisons.

---

### 🚀 Features

* ✅ Simple and clear implementation
* ✅ No third-party libraries
* 🧠 Ready for future algorithm additions (KMP, etc.)
* 📊 Benchmark and performance tracking supported

---

### 🧮 Algorithm 1: Naive Search

**Description:**
The naive search algorithm checks for the substring match by iterating through each position of the `haystack` and comparing it with the `needle`.

**Time Complexity:** `O(n * m)`
**Space Complexity:** `O(1)`

Where:

* `n` → length of `haystack`
* `m` → length of `needle`

---

### 📊 Performance

You can visualize your time-performance results here:

```markdown
![Performance Comparison](./images/performance.png)
```

*(The image file should be located at `./images/performance.png`.)*

---
