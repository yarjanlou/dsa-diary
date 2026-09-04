# 01 — Hashing & Frequency Counting

## Concept

Hashing & Frequency Counting is the pattern of using a hash-based structure (a set or a map) to remember information about elements you've already seen, so you can answer questions about the data in less time than a brute-force scan would take.

## Why It Matters

A huge number of interview problems reduce to some version of "have I seen this before?" or "how many times does this appear?". Recognizing that reduction turns an O(n²) nested-loop solution into an O(n) one, and it's one of the highest-leverage patterns to have automatic, because it shows up as a building block inside much harder problems too.

## Recognition Signals

- The brute-force approach involves comparing every element to every other element (nested loops).
- The problem asks about duplicates, uniqueness, or "have I seen this before".
- The problem asks about counts, frequencies, or "most/least common".
- The problem involves pairs, complements, or things that must sum/match to a target.
- Elements need to be grouped by some derived property (anagram, remainder, category).

## Core Mental Model

- **Fast lookup**: a hash map/set gives average O(1) lookup instead of O(n) scanning.
- **Membership checking**: a set answers "is this in the collection?" without caring how many times.
- **Remembering previously seen information**: as you iterate once, the map/set becomes your memory of the past.
- **Frequency counting**: a map from value → count turns "how many times" into a single pass.
- **Trading space for time**: you accept O(n) extra memory to collapse O(n²) time down to O(n).

## Common Patterns

- **Seen Set** — track whether an element has already been encountered (duplicates, cycles).
- **Frequency Map** — count occurrences of each element (anagrams, majority element, mode).
- **Complement Lookup** — for each element, check whether the value needed to complete a pair/target already exists in the map (classic Two Sum).
- **Grouping by Key** — bucket elements under a derived key, e.g. a sorted string or a normalized form (Group Anagrams).
- **Lookup to avoid nested loops** — anytime an inner loop exists only to search for a match, replace it with a map/set lookup.

## Common Mistakes

- Reaching for a nested loop out of habit before checking whether a set/map lookup would do the same job in one pass.
- Forgetting to handle counts correctly — checking `if value in map` when you actually need to increment/compare a count.
- Using a list where a set was intended, silently turning an O(1) lookup back into O(n).
- Not considering hash collisions/key design for non-primitive keys (e.g. using an array or object as a map key without normalizing it, like sorting characters for anagrams).
- Mutating a map/set while iterating over it.

## Problem Progression

| # | Problem | Difficulty | What it teaches | Status |
|---|---|---|---|---|
| 1 | Contains Duplicate | Easy | Seen Set — basic membership checking | ⏳ Not Started |
| 2 | Valid Anagram | Easy | Frequency Map — comparing counts between two collections | ⏳ Not Started |
| 3 | Two Sum | Easy | Complement Lookup — replacing a nested loop with a single pass | ✅ Solved |
| 4 | Group Anagrams | Medium | Grouping by Key — bucketing elements by a derived key | ⏳ Not Started |
| 5 | Top K Frequent Elements | Medium | Frequency counting combined with sorting/selection | ⏳ Not Started |

## Mastery Check

Before considering this concept mastered, we should be able to:

- Spot, within a minute of reading a problem, whether it reduces to a seen-set or frequency-count problem.
- Explain why a hash map/set turns an O(n²) approach into O(n), including the space tradeoff.
- Choose correctly between a set (membership only) and a map (membership + count/value).
- Design a normalized key for grouping problems (e.g. sorted string, sorted tuple) without hesitation.
- Implement any of the five patterns above from scratch without referencing notes.
