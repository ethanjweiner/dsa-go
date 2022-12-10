# Greedy Stock Predictions

## Nieve Approach (O^N^3)

Given an array of `prices`:
- For each price `a`: (`O(N)`)
  - Iterate through subsequent prices `b`
    - Upon finding first subsequent price `b > a`
      - Iterate through subsequent prices `c`
        - Upon finding first subsequent price `c > b`
          - Return `true`

## Optimization

1. Current time complexity: `O(N^3)`
2. Best imaginable: Might need to examine every element -> `O(N^3)`
3. Greedy approach: Continually grab what we THINK are the highest/lowest points
  - Once we find a trend that works, we are done

Given an array of `prices`:
- Initialize `lowest` to `prices[0]`
- Intialize `middle` to `Infinity`
- Iterate through `prices`:
  - If `price` is lower than `lowest` -> Reassign `lowest` (best lowest so far)
  - If `price` is > `lowest` and `<` middle -> Reassign `middle`
  - If `price` is > `middle` so far -> Upward trend is found