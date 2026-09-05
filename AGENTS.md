# AGENTS.md — Repository Architecture & Agent Instructions

This file is the canonical reference for how this repository (`dsa-diary`) is organized and how AI coding agents (Claude Code, Claude in VS Code, OpenAI Codex, or any other repository-instruction-aware agent) must work with it.

It does not document every file. It documents the **rules an agent must follow** before adding, moving, modifying, or organizing LeetCode problems, so that months from now an agent can correctly answer:

> "Where does this new problem belong, what files should I create, what metadata should I add, and what should I avoid changing?"

If future user instructions conflict with this document, follow the more recent explicit instruction — and if the architecture is intentionally changed, update this file so future agents use the new architecture. Do not treat this file as a reason to refuse a requested change.

---

## Learning methodology

This repository is a structured DSA learning diary, not a flat list of solved problems. The methodology it reflects:

```text
Choose concept
    ↓
Learn concept
    ↓
Understand recognition signals and mental models
    ↓
Solve a deliberate progression of problems
    ↓
Review and reflect
    ↓
Master the concept
    ↓
Move to the next concept
```

Curriculum decisions — which concept comes next, which problems to solve, how many problems a concept requires, whether a concept is mastered — belong to the human's DSA learning process. **Agents must faithfully represent those decisions in the filesystem and documentation, not invent or redesign them.**

---

## Repository structure

```text
LeetCode/
├── 00-legacy-warmups/
│   ├── README.md
│   ├── <existing legacy problem>/
│   ├── <existing legacy problem>/
│   └── ...
│
├── 01-hashing-and-frequency-counting/
│   ├── README.md
│   ├── 01-<problem-slug>/
│   │   ├── README.md
│   │   └── solutions/
│   │       ├── README.md
│   │       ├── main.js
│   │       ├── main.go
│   │       └── ...
│   ├── 02-<problem-slug>/
│   └── ...
│
├── 02-<next-concept>/
│   ├── README.md
│   └── ...
│
└── ...
```

Concept directories use an **append-only numeric prefix** (`01-`, `02-`, `03-`, ...) representing the order the concept was studied — it is **not** a permanent importance ranking. Never renumber existing concepts because the curriculum changes.

As of this writing, the repository contains:
- `LeetCode/00-legacy-warmups/` — 12 pre-curriculum problems.
- `LeetCode/01-hashing-and-frequency-counting/` — the first structured module, containing only `README.md` (its problem progression is planned but no problem folders exist yet — they are created only as each problem is actually solved).

---

## Legacy problems (`00-legacy-warmups/`)

Problems solved before the structured curriculum belong here permanently.

- Do **not** retroactively reorganize legacy problems into concept directories, even if their pattern is now obvious.
- Do **not** move a legacy problem unless explicitly instructed — conceptual fit does not change physical location.
- Future concept READMEs may reference relevant legacy problems by link; that's the extent of the connection.

---

## Concept modules (`LeetCode/NN-concept-name/`)

Each concept has exactly one directory containing a `README.md` and the problem directories primarily belonging to it.

The concept README is a **concise, durable summary**, not a transcript of a teaching conversation. Detailed explanation can live in chat history; the repo holds the durable knowledge worth revisiting later. Typical sections:

- Concept
- Why It Matters
- Recognition Signals
- Core Mental Model
- Common Patterns
- Common Mistakes
- Problem Progression (table)
- Mastery Check
- What We Learned (filled in after the concept is worked through)
- Relevant legacy problems (optional links)

---

## Problem organization

**Each problem exists in exactly one physical location**, determined by its **primary concept**. Example:

```text
LeetCode/01-hashing-and-frequency-counting/01-two-sum/
```

A problem may reinforce multiple concepts, but it is never duplicated. Do not create parallel copies like:

```text
hashing/two-sum/
sliding-window/two-sum/
```

for the same problem, and do not use symlinks to fake a second location. Secondary concepts are represented as metadata/text in the problem's own README (see below), not as additional directories.

### Problem directory structure

```text
<problem-slug>/
├── README.md
└── solutions/
    ├── README.md
    ├── main.js
    ├── main.go
    ├── main_<approach>.js
    ├── main_<approach>.go
    └── performance/
        └── ...
```

Not every problem needs every file (e.g. no `performance/` if no benchmark exists). Do not create empty directories or placeholder files "just in case."

### Problem README

Contains the problem writeup plus lightweight metadata, e.g.:

```markdown
| Field | Value |
|---|---|
| LeetCode # | 1 |
| Difficulty | Easy |
| Primary concept | Hashing & Frequency Counting |
| Secondary concepts | — |
| Progression order | 1 / 5 |
| Status | ✅ Solved |
```

No JSON, YAML, databases, or sidecar metadata files unless explicitly requested. Keep the README focused on the problem itself.

### Solutions

Live under `<problem>/solutions/`. Preserve existing conventions:

- Primary solution: `main.js`, `main.go`.
- Multiple approaches: `main_<approach>.js` / `main_<approach>.go` (underscore suffix — e.g. `main_naive.js`, `main_optimal.js`, `main_kmp.js`).
- Preserve existing benchmark/performance artifacts (screenshots, images) — never delete them.
- Do not rewrite existing solutions for cosmetic naming reasons unless explicitly asked. Renaming to fix a naming inconsistency (e.g. hyphen → underscore) is acceptable when it's unambiguous, loses no history, and is done via a Git-aware move — but broad cosmetic renaming is not something to do proactively.

---

## Adding a NEW problem (structured curriculum)

1. Determine which concept module it belongs to (primary concept).
2. Determine its position in that concept's progression.
3. Add it to that concept's progression table in the concept README.
4. Create exactly one problem directory under the appropriate concept, named with its progression order + LeetCode URL slug (e.g. `01-two-sum`).
5. Create the problem `README.md`.
6. Create the `solutions/` directory only when actually solving the problem (not before).
7. Add solution files as appropriate.
8. Update the problem's own metadata table.
9. Update the concept README's status/progression.
10. Update the root README's progress information when appropriate.

Never duplicate the problem into a second location. Never place a new curriculum problem inside `00-legacy-warmups/`. Never create a new concept directory just because one problem touches a secondary pattern.

---

## Adding a NEW concept

```text
LeetCode/NN-concept-name/
└── README.md
```

- Use the next available concept number; never renumber old concepts.
- The module README should define: what the concept is, why it matters, recognition signals, important patterns/variations, problem progression, and mastery/checkpoint criteria.
- Problem directories are added only as the progression is actually worked through — do not pre-create empty placeholder problem folders.

---

## Cross-concept problems

Represent secondary concepts with metadata/Markdown, never a second physical directory:

```markdown
| Primary concept | Sliding Window |
| Secondary concepts | Hashing & Frequency Counting |
```

A relevant concept README may add a concise cross-reference link back to the problem. Keep this lightweight — no tagging system, no index files.

---

## Naming conventions

- Concept directories: `NN-kebab-case`
- Problem directories: progression order + LeetCode URL slug, kebab-case (e.g. `01-two-sum`)
- Markdown files: `README.md`
- Primary JS solution: `main.js`
- Primary Go solution: `main.go`
- Additional approaches: `main_<approach>.<ext>` (underscore)

Avoid arbitrary naming changes. Preserve established naming in existing work unless there's an explicit reason to migrate it.

---

## What agents should NOT do

Unless explicitly instructed, do not:

- Reorganize the architecture.
- Renumber concepts.
- Retroactively move legacy problems.
- Duplicate problems across concepts.
- Create symlinks.
- Create JSON/YAML tracking systems.
- Create unnecessary notes/mastery/progress files.
- Create placeholder folders for unsolved problems.
- Rewrite existing solutions.
- Delete existing benchmarks or screenshots.
- Replace the repository's learning methodology with a generic LeetCode list.
- Add documentation for its own sake.
- Commit or push changes unless explicitly asked.

---

## Decision-making rules

When unsure where something belongs:

1. Prefer the existing architecture.
2. Prefer one canonical physical location.
3. Prefer primary concept + secondary metadata over duplication.
4. Preserve existing history and work.
5. Prefer simple Markdown over additional tooling.
6. Do not invent new organizational systems without a clear need.
7. Ask for clarification only when the decision materially affects the repository architecture; otherwise make the smallest reasonable change.

---

## Safe modification principle

1. Inspect before modifying.
2. Preserve existing work.
3. Make the smallest change that satisfies the request.
4. Keep the structure consistent with this document.
5. Validate the resulting tree and relevant files.
6. Report what changed.

Prefer Git-aware operations (`git mv`) for moves or renames so history is preserved.
