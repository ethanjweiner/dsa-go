# Lowest Common Ancestor

## Problem

- Input: A `root`, and two nodes `p` and `q`
- Output: The LCA of `p` and `q`
- Note: The LCA may be either `p` or `q`, because a node can be a descendant/ancestor of itself (by definition)

## Questions

- Can `p` and `q` be the same node? (assume no)
- Can I use additional data structures to solve the problem?

## Ideas

- Record a trail of ancestors, as you traverse. Once you find the node, complete. Then, compare the two nodes, and find the furthest ancestors.
- Continually try out different roots, if the root contains both descendants -> it is a candidate
- Keep trying until the root no longer contains both descendants

## Data Structures

- Storing ancestors:
  - In a slice?
  - In a linked list (prepend newest ancesetors?)

## Algorithm

### Ancestors trail

Given a `root`, `p`, and `q`:
- Compute the ancestors of `p` (lowest first)##
- Compute the ancestors of `q` (lowest first)##
- Find the common ancestor which is closest to the start of both `p` and `q`

- Time Complexity: `O(log n)` (binary tree search)
- Space Complexity: `O(log n)` (create 2 linked lists of ancestors)

### Try out multiple ancestors

Given a `root`, `p`, and `q`:
- If `root` is `nil`, it can not be a common ancestor -> return `nil`
- Try finding the LCA of `p` and `q` on the left. If found, return it.
- Try finding the LCA of `p` and `q` on the right. If found, return it.
- Else:
  - Search for `p` and `q` on the current tree
  - If both found, return `root`
  - Otherwise, return `nil` (no LCA)

*Main Idea*: Start with the smallest ancestors, & work up

### BST-Specific solution

Given a `root`, `p`, and `q`:
- If `p` & `q` < root -> search left (ancestor must be on left)
- If `p` & `q` > root -> search right (ancestor must be on right)
- Else, the `root` must be the ancestor