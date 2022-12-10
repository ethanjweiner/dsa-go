# Coin Game

Given a handful of `n` coins, and the `currentPlayer` determine the winner (`you` or `them`) of the game (assuming both players play optimally).

Rule: Each player must remove 1 or 2 coins on each turn, player to remove final coin loses

Small examples:
- 1 coin, you -> you lose
- 2 coins, you -> you win
- 3 coins, you -> you win
- 4 coins, you -> you lose
- 5 coins, you -> sub 1 coin -> you win
- 6 coins, you -> sub 2 coins -> you win
- 7 coins, you lose
- 8 coins, you -> sub 1 coin -> you win
...

Pattern: Pattern: `1 + 3x, x is an integer` -> you lose, otherwise you win

Winning options:
- Base case = 2 coins: `2 + N` where `N` is all multiples of `1`: 

## Traditional Approach

Given a handful of `n` coins, and the `currentPlayer`
- If 1 coin -> return other player
- If 2 or 3 coins -> return current player
- Else:
  - If `n - 1` coins guarantees the next player loses, or `n - 2` coins guarantees the next player loses -> return `currentPlayer`
  - Otherwise, return other player

### Efficiency

Given `n` coins:
- 2 recursive calls per 

Recursivee call structure (worst case):
                  n
                /   \
            n - 1  n - 2
          /   \       /   \
      n - 2  n - 3  n - 3  n - 4
               ......
              1 | 2 | 3

`O(2^n)` pattern

## Improvements

1. Current time complexity: `O(2^n)`
2. Best imaginable:
  - `O(n)` (memoization -> entire half of recursive tree structure is cut off -> linear instead -> `O(n)`)
  - `O(1)`: Use a numerical property of `n` to determine winner
3. If `n == 3N + 1, N is an integer` -> you lose, else you win
  - Put another way: If `(n - 1 % 3) == 0` -> you lose, else you win 