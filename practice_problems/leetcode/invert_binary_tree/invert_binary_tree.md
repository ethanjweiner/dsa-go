# Invest Binary Tree

## Problem

- Input: Binary tree
- Output: The same root inverted recursively

## Algorithm

### Top-Down, Recursive

Given a `root`
- If `root` is `nil` (empty) -> return
- Invert the left
- Invert the right
- Swap the root's left & right nodes
- Return