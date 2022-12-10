# Maximum Sum Subarray

## Problem

- Input: Array of positive `numbers` and pos num `k`
- Output: Max sum of any contiguous subarray of size `k`

## Algorithm (Sliding Window)

Given `numbers` and `k`:
- Initialize a `maxSum` of the first `k` numbers:
  - Iterate from `0` to `k - 1`, determining the sum
- Initialize a `currentSum` to the `maxSum`
- (For each window) Iterate from a `windowStart` of 1 to length of `numbers` - k (inclusive):
  - Subtract the element at `numbers[windowStart - 1]` from `currentSum`
  - Inrement `currentSum` by `numbers[windowStart + k - 1]`
  - If the `currentSum` > `maxSum` -> reassign `maxSum`
- Return `maxSum`