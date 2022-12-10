# Stock Prediction

Given a set of `predictions`, compute the greatest profit that could be made from a single buy/sell combo (buy first, sell second).
(i.e. max (|low - high|)

Best imaginable time complexity: `O(N)`

*Idea*: Iterate until the first increase -> buy, iterate until the first decrease -> sell
*Greediness*: Find the lowest possible transaction, moving from the left, & the highest possible transaction, moving from the right
*Idea 2*: Continually close in from both sides

Pattern recognition:

3 7 2 7 1 8 4
B S B S B S B
8 - 1 = 7

5 9 4 128 2 54 67 50 129 1
L H    H  L2             L3        
H

Idea: Find all the low-high gaps, and store the max

## Greedy Approach

*Main Idea*: Always finding the profit from only the **lowest buy price** so far, disregarding any future higher buy prices (irrelevant)

- Initialize a `buyPrice` to the first #
- Initialize a `bestProfit` to `0`
- Iterate over the `predictions`:
  - If `prediction` < `buyPrice` -> reassign new lowest ` buyPrice`
  - Else:
    - Compute the `profit`: `prediction - buyPrice`
    - If `profit` > `bestProfit` -> reassign `bestProfit`

*Efficiency*: `O(N)` (iterates through predictions once)