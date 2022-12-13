# Subarrays With Product

- Input: Array of positive `numbers` and a `target` number
- Ouptut: The contiguous subarrays in `numbers` whose product < `target`

## Ideas

### Traditional Approach

Nested iteration: Iterate up until the end. Using a nested loop, continually build contiguous subarrays (one longer each time). Append a copy of that subarray if its product < target. Break from nested loop once a subarray's product > target (avoids further unnecessary computation).

Current time complexity (worst case): `N + (N-1) + ... + 1` subarray creations copying # of subarray elements each time -> `O(N^3)`

### Best Imaginable Time Complexity

`O(N^2)`: We might have `N + (N-1) + ... + 1` subarrays (in the worst case)

### Two Pointers + Sliding Window Approach

Set a current product to `1`. Set two pointers, initially both set to the first #. Continue to shift over the right pointer until the product is greater than or equal to the target. Generate all contiguous subsets starting from the left pointer and ending at the right pointer (exclusive). At this point, shift over the left pointer to the right until the product is less than or equal to the target (dividing to adjust the product). Repeat by shifting the right pointer.

[2, 5, 3, 10]
l   r

subsets generated:
[2], [5], [5, 2], [3], [3, 5], [10]

### Time & Space Complexity

Time: `O(n^3)`
Space: `O(n)` for the intermediate list, `O(n^2)` for the output list (`n + (n - 1) + ... + 2 + 1` total subarrays)