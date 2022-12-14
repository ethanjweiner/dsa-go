# Path with Given Sequence

## Problem

- Input: Number binary tree `bt` and `sequence` of numbers
- Output: Boolean indicating whether the sequence is present as a root-to-leaf path

## Ideas

No-accumulator: Pass a new version of the `sequence` to find each time.
W/ accumulator: Pass the sequence created so far, once the `sequence` begins to diverge -> return false. 

## Algorithm (no accumulator)

Very DECLARATIVE

Given a `bt` and a `sequence`:
- [Base case] If leaf:
  - Return true if leaf's data equals only value in `sequence`
- Return true if:
  - Current data is first value AND:
  - left comprises remaining sequence OR right comprises remaining sequence

### Time/Space Complexity

Time: `O(N)` (visit every node)
Space (`N` = binary tree nodes, `K` = sequence size):
- Worst case = `O(N)` (call stack might string all nodes)
- Average case (balanced tree) = `O(log N)` (reslicing doesn't take up extra)

## Algorithm (accumulator)

Tree:              Sequence    ??? INDEX
                   
     1             [1, 3, 6]   0
   /   \            
  2     3          [1, 3, 6]   1
 / \   / \ 
4  5  6   7        [1, 3, 6]   2

Imagine:
- `sequence` is fixed on each iteration
- Then, how do you 
