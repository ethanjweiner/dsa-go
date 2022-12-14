# Quadruple Sum to Target

- Input: An array of `numbers` and a `target` number
- Output: Sum of triplet whose sum is closest to `target`
- If 2 triplets are as close to target -> return sum of smallest triplet

## Ideas

### Traditional Approach

Find all triplets (`C(n, 3)`) and for each triplet, analyze its sum.
Would probably use a nested loop to generate triplets, iterate through all triplets, and set a max.
Time Complexity: `O(3 * C(n, 3))`

### Best Imaginable Time Complexity

- Not necessary to analyze every triplet
- However, each # nmust be analyzed at least once
- Best time complexity: `O(N)`

### Two Pointers + Sliding Window Approach

Main Idea: We have three pointers: `X`, `Y` and `Z`. By having a *sorted* array, there will be a point at which further checking is no longer effective. We want to stop at that point.

For each iteration of moving `X` to the right:
- Try to find a `Y` and `Z` that sum as close to the remaining difference as possible. Once that difference is found:
  - If it is better than the `bestSum`, set it
  - If it is equal to the `bestSum`, take the lower of the two sums
- Move `X` to the right and repeat, up until:
  - `X` is at the 3rd to last index
  - `X - target` > `|bestSum - target|` (no point in going further)

### Sub-problem: Finding the best pair

-2 0 1 2 3, 5, 8 target=2
X  Y     Z          

Idea w/ pair:
- Keep moving Z right until incrementing any more => further deviations.
- Then, do the same w/ Y

For each iteration of moving `Y` to the right:
- Try to find a `Z` that is as close to the remaining difference as possible. Keep traveling right until the distance to the sum is no longer improving, but worsening

X         | -2 ....
Y + Z     | ~4 ....       
sum       |               1 2 3 8(x)
Y         |               0    
Z         |               1 2 3 8

Greedy approach: just skip over the #s that will make the difference worse.

## Example

### Algorithm 1

target=1        bestSum        
[-3, -1, 1, 2], Infinity
X     Y  Z      -3
            Z   -2
X        Y  Z   0
      X  Y  Z   2X 0
Return

target=3           bestSum        
[-3, -1, 1, 2, 4], Infinity
X     Y  Z         -3
            Z      -2
               Z   0
               ...
X           Z  Z   3 -> done (return 3)

### Algorithm 2

Goal: For each X, find the optimal pair (starting at polar ends)

target=3           bestSum    remaining  pairSum
[-3, -1, 1, 3, 4], Infinity
 X    Y        Z              6          3    
         Y                    6          5
                              6          6
## Algorithm 1

Given an array of unsorted `numbers` and a `target`:
- Produce a `sortedNumbers` version of `numbers`
- Set a `bestSum` to `Infinity`
- For each `X` from `0` to length of `numbers` - 3 (inclusive):
  - Find the `bestRemaingSum` from `Y` and `Z` (in a slice from index `X + 1` to the end) ST `Y` and `Z` are closest to `target - X`##
  - Compute `currentSum = X + bestRemainingSum`
  - If `|currentSum - target| < |bestSum - target|` -> reassign `bestSum` to `currentSum`
  - **Optimization (implement later)**: Return once we know we have passed the best sum
    - Condition: We have surpassed the sum and its distance from the target >= previous
    - Or, in computation terms: `currentSum - target >= bestSum - target`
    - Since the array is sorted, no further processing would lead to a better sum
- Return `bestSum`

### Find pair closest to a target sum (two sum problem)

Given an array of `sortedNumbers` and a `target`:
- Set a `bestSum` to `Infinity`
- For each `Y` from `0` to length of `numbers` - 2 (inclusive):
  - Find the best number `Z` (in a slice from index `Y + 1` to the end) ST `Z` is closest to `target - Y`##
  - Compute `currentSum = Y + Z`
  - If `|currentSum - target| < |bestSum - target|` -> reassign `bestSum` to `currentSum`
  - **Same optimizaiton**
- Return `bestSum`

### Find number closest to a target

- Set a `bestNum` to `Infinity`
- For each `Z` from `0` to the length of `numbers` - 1 (inclusive)
  - If `|Z - target| < |bestNum - target|` -> reassign `bestNum` to `Z`
  - **Same optimization**: If `Z - target >= |bestNum - target|` -> break early
- Return `bestNum`

## Time Complexity

### Algorithm 1

Sorting -> `O(n log n)`
Three nested iterations to move pointers -> `O(n^3)`
Overall time complexity: `O(n^3 + nlogn)` = `O(n^3)` (doesn't beat traditional)

### Algorithm 2

Sorting -> `O(n log n)`

Moving X (first pointeer) = one iteration (max `n` steps)
Finding pair = second nested iteration (max `n - 1` steps)
-> `O(n^2)`

Overall time commplexity: `O(n^2 + nlogn)` = `O(n^2)`