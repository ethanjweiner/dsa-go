# Two Sum Problem

Given an array of `numbers` and an integer `n`, determine whether any two `numbers` add to `n`

## Traditional Approach

For each `num1` in `numbers`:
- For each `num2` in `numbers` (except `num1`):
  - If `num1` & `num2` add to `n` -> return `true`

## Efficiency

1. Current efficiency: `O(N^2)` worst case
2. Best possible: `O(N)` (for each #, check whether `n - num` exists in `numbers`)
3. Optimization technique: Convert `numbers` to hash table (`O(1)` lookup)

- Transform `numbers` into hash table (`O(N)`)
- For each `num` in `numbers`: (linear time) (`O(N)`)
  - Determine the other desired `addend`: `n - num`
  - If `addend` is in `numbers` -> Return `true` (constant time)

Time complexity = `O(N + N)` = `O(2N)` = `O(N)`