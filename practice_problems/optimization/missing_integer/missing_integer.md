# Missing Integer

## Using a hash table

Given an array of `integers`:
- Turn array -> hash table (set)
- Iterate from 0 upwards:
  - The first # not in the hash table is the missing integer
  - Return that #

## Pattern-based approach

Compute the sum of the array, utilize sum to determine missing #

e.g. [1, 3] -> Sum = 4, missing = 2
e.g. [, 2, 3] -> Sum = 5, missing = 1

Generate the sum up to the max # + generate the ACTUAL sum -> determine difference

Given an array of `integers`:
- Initialize an `actualSum` to zero
- Initailize a `max` to the first integer
- Iterate over `integers`: `O(n)`
  - If the current integer > `max` -> Reassign
  - Add to the actual sum
- Iterate over `0..max` and compute `desiredSum`
- Return `desiredSum - actualSum`