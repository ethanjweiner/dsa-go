# Minimum Depth of Binary Tree

## Problem

Find the minimum depth of a binary tree. The minimum depth is the number of nodes along the shortest path from the root node to the nearest leaf node.

- Input: A tree `bt`
- Output: The number of nodes along the shortest path from the root node to any leaf node

## Ideas

Ask group: How would you solve this problem? (code-independent)

## Tree

      12
     /  \
    7    1
   /     / \
  9     10  5
         \
          11

## DFS

Main Idea: Search down every path, try to find the lowest.

### w/o Accumulator (Declarative)

Before we start thinking about the problem, think about the data structure. We have a tree -> case analysis:

Main Idea: 1 + the min(left min, right min)

Given a `bt`:
- If `bt` is a leaf, then return `1`
- If `bt` left is `nil` -> return `1` + right subtree min depth
- If `bt` right is `nil` -> return `1` + right subtree min depth
- Else (both exist), return the minimum of the two

### w/ Accumulator (Imperative/Procedular)

Ask yourself: What could I use at each previous stage?

### Time Complexity

If `N` is the total # of nodes:

Worst case:
Average case:
Best case:

### Space Complexity

Worst case:
Average case:
Best case:

## BFS

Main Idea: Search each level, once we reach a leaf in any level, return the depth.

Walk through using a queue

Given a `bt`:
- Initialize a `depthSoFar` to `1`
- Initialize a `queue` to the root node
- Loop:
  - Initialize a `nextQueue` to an empty queue
  - Until the `queue` is empty:
    - Dequeue the next element
    - If that element is a leaf, immediately return the `depthSoFar`
    - Enqueue that next element's left & right -> `nextQueue`
  - Increment the `depthSoFar` by 1

### Time Complexity

If `N` is the total # of nodes:

Worst case:
Average case:
Best case:

### Space Complexity

If `N` is the total # of nodes:

Worst case:
Average case:
Best case:

## [If time] Implement DFS w/ a stack

## Logistics

Bunch everything into Tues meeting?