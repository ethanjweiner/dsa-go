# Binary Tree Path Sum

## Problem

- Input: `tree` and number `S`
- Output: Boolean indicating whether `tree` has a sum of node values equaling `S`
- Question: Can data only be positive integers?
  - If so, can optimize to return once sum to compute becomes negative

## Ideas

Traverse DFS the tree with an accumulator. Once we reach a leaf, test whether the accumulator (sum) equals S. If so, return true. Otherwise, keep traversing.

Can all DFS algorithms use a simple function for traversal? (trick)

## Algorithm (Non-Cumulative)

Given a `bt`, a `sum` to find
- [Base case] If leaf equals `S` -> return true
- Initialize a `nextSum` to `S - data`
- Return true if `left`'s sum equals `nextSum` (recursive) or if `right`'sum sum equals `nextSum`)
 - Return true if `right`'s sum equals `nextSum` (recursive)

## Time Complexity

Given a binary tree with `n` nodes, the time complexity is `O(n)` (might have to visit every node).

## Space Commplexity

The space complexity is `O(n)` in the worst case (unbalanced tree) <- recursive calls on call stack. In the average case, it's `O(log n)`, because the tree has a height.