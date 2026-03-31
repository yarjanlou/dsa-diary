# Perfect Square Checker in Go

This project demonstrates two different approaches to determine whether a given number is a **perfect square** in Go:

1. Binary Search (Iterative Approach)
2. Newton’s Method (Optimized Mathematical Approach)

---

## 📌 Problem Definition

A number is a **perfect square** if it can be expressed as:

```
x * x = num
```

Examples:

* 16 → ✅ (4 × 4)
* 25 → ✅ (5 × 5)
* 20 → ❌

---

## 🚀 Solution 1: Binary Search

### 🔧 Implementation

```go
func isPerfectSquare(num int) bool {

	if num == 1 {
		return true
	}

	var start int = 1
	var end int = num
	var mid int = num / 2

	for start < mid {
		if mid * mid == num {
			return true
		}

		if mid * mid > num {
			end = mid
		}

		if mid * mid < num {
			start = mid
		}

		mid = (start + end) / 2
	}
	
	return  false
}
```

---

### 🧠 How It Works

* The algorithm searches for a number `mid` such that `mid * mid == num`.
* It repeatedly halves the search space:

  * If `mid²` is too large → search left
  * If `mid²` is too small → search right

---

### ⚠️ Limitations

* ❌ Risk of infinite loop (`start = mid` / `end = mid`)
* ❌ Incorrect loop condition (`start < mid`)
* ❌ Possible integer overflow (`mid * mid`)
* ❌ Not handling `0`

---

### ⏱ Complexity

* Time: **O(log n)**
* Space: **O(1)**

---

## ⚡ Solution 2: Newton’s Method

### 🔧 Implementation

```go
func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}

	x := num

	for x > num/x {
		x = (x + num/x) / 2
	}

	return x*x == num
}
```

---

### 🧠 How It Works

Newton’s method approximates the square root using:

```
x = (x + num/x) / 2
```

* Starts with an initial guess (`x = num`)
* Iteratively improves the estimate
* Stops when `x` converges to √num

---

### ✅ Advantages

* ✅ Faster convergence than binary search
* ✅ Fewer iterations (usually < 10)
* ✅ Avoids overflow using `num/x`

---

### ⚠️ Considerations

* Final validation (`x*x == num`) is still required
* Uses division (slightly more expensive than addition/multiplication)

---

### ⏱ Complexity

* Time: **O(log n)** (but faster in practice)
* Space: **O(1)**

---

## 🆚 Comparison

| Feature       | Binary Search   | Newton’s Method   |
| ------------- | --------------- | ----------------- |
| Simplicity    | Easy            | Moderate          |
| Speed         | Good            | Very Fast         |
| Iterations    | More            | Fewer             |
| Overflow Risk | Yes (`mid*mid`) | Avoidable         |
| Accuracy      | Exact           | Needs final check |

---

## ✅ Recommendation

* Use **Newton’s Method** for performance-critical scenarios
* Use **Binary Search** for simplicity and readability (after fixing edge cases)

---

## 🛠 Example Usage

```go
func main() {
	fmt.Println(isPerfectSquare(16)) // true
	fmt.Println(isPerfectSquare(20)) // false
}
```

---

## 📌 Final Note

If you use the binary search approach in production, make sure to:

* Fix loop conditions
* Avoid overflow
* Handle edge cases properly

Newton’s method is generally the **preferred solution** for this problem.

---
