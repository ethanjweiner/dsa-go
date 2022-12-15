# Binary Tree Level Order Traversal

## Problem

- Input: A binary tree `bt`
- Output: An array of subarrays representing its `levels`

## Ideas

- Start at the top-level node
- Initialize a queue, fill up the queue
- Once the queue 

When to start emptying the queue? (once you finish adding all the elements)

What is the "stopping point"? Once the queue/level is empty?

                Queue         New Queue    Levels 
     1          1             2, 3         [[1]]
   /   \        | \
  2     3       2, 3          ...
 / \   / \     
4  5  6   7     4, 5, 6, 7    Empty -> reassign means we finish


## Algorithm

Given a `bt`:
- Initialize a `queue` w/ just the root node's data
- Initialize `levels` to an empty slice of int slices
- While the `queue` is non-empty:
  - Initialize a `newQueue`
  - Dequeue every element from `queue`:
    - Add to a new `level`, append to `levels`
    - Queue the left & right to a `newQueue`
  - Reassign the `queue` to the `newQueue`
- Return the `levels`