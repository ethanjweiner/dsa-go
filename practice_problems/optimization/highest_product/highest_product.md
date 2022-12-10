# Highest Product

Given an array of `numbers` (including negatives), find the highest product of any two numbers.

1. Current efficiency: `O(N^2)`
2. Ideal efficiency: `O(N)`

## Ideas

- Collect TWO minimums (greatest negative) & maximums
- Return the greater of the two products

Maxs: [3, 4]
Mins: [-5, 0]

1  2  3  -5  4  -3  -6  3  10  8  2  -5  3
X  X  X  X   X ....  

Given an array of `numbers`:
- Initialize 2 slices: `mins` and `maxs` (to array of zeroes)
- Iterate over `numbers`:
  - If `number` is positive, and greater than either element in `maxs`: (3 comparisons)
    - Replace the lesser element in `maxs` (2 comparisons)
  - Do the same for negatives
- Compute the product of `mins` and `maxs`, and return the greater product
