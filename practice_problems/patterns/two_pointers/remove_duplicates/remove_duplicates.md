# Remove Duplicates

## Problem

- Input: Array (slice) of sorted `sortedNumbers`
- Output: The new length of the array
- Side Effects: All duplicates should be removed from the array (in-place deletion)
- Space complexity should be `O(N)`

## Ideas

### Traditional Approach

[Space complexity of `O(N)`] Store previously seen elements in an array or hash table -> add to new array only if not already seen

### Two Pointers Approach

Initialize two pointers to 0 & 1 (indices next to each other). If they have equal values, remove the value at the second pointer from the array. Continue until the second pointer becomes out of bounds.

[Guard clause] If the length of the array <= 1 -> return

## Algorithm
### Algorithm 1

Given a slice of `sortedNumbers`:
- [Guard clause] If the length of the array <= 1 -> return
- Initialize `p1` and `p2` to 0 & 1
- Loop while `p2` < length of `sortedNumbers`
  - If `sortedNumbers[p1] == sortedNumbers[p2]`:
    - Delete the element at `p2`
  - Else, increment `p1` and `p2` by 1
- Return the length of the now-unique `sortedNumbers`

### Algorithm 2

Given a slice of `sortedNumbers`
- Initialize a `nextNonDuplicate` to 1
- Iterate over the `sortedNumbers` (starting at 1) tracking the current `idx`
  - If `sortedNumbers[idx]` does not equal `sortedNumbers[nextNonDuplicate - 1]`:
    - Set `sortedNumbers[nextNonDuplicate] = sortedNumbers[idx]`
    - Increment the `nextNonDuplicate` 
- Return a new slice from 0 to the final `nextNonDuplicate` (plus 1)

### Element deletion (w/ order preservation)

Given a `slice` and an `indexToDelete`:
- [Guard clause] If `indexToDelete` >= length -> return error
- `append` the elements up to `indexToDelete` with the elements after `indexToDelete`
- Return the appended slice

## Time Complexity

- `n` analyses of values at pointers
  - [Worst case] For each analysis, delete an element (`O(n)` operation <- array must be moved)

## Implementation

- Deleting element from a slice

## Future Optimizations

- Instead of deleting elements progressively, simply reassign elements and return a sliced version