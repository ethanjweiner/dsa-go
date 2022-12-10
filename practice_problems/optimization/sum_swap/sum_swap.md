# Sum Swap

`O(N + M)` algorithm:

Given `arr1` and `arr2`:
- Convert `arr2` to hash table `O(M)`
- Compute `shiftAmount` (amount each sum must change by to meet the middle): `sum(arr1) - sum(arr2) / 2`
- For each `integer` in `arr1`:
  - Determine the integer in `arr2` that would produce shift amount: `integer + shiftAmount`
  - If that value is an integer & `arr2` hash contains it, then return `true`
