# Largest Subsection Sum

Given a an array of `numbers`, return the largest sum

## Example

[3, -4, 4, -3, 5, -9]

## Traditional Approach (Generate all subsections)

For each element `i` in `array`: 
- Track a `subsectionSum`
- Iterate from `i` to the end:
  - Each time, add 2nd element to `subsectionSum` 
  - If `subsectionSum` > current largest sum -> add to it (1 comparison)

`N + N - 1 + ... + 2 + 1` steps = `O(N^2)`

## Optimization

1. Current efficiency: `O(N^2)`
2. Best imaginable: `O(N)` (need to inspect each # at least once)
3. Pattern recognition: Whenever the leading sum becomes neegative, then it is *guaranteed* that there is a better option -> skip over

Track a `currentMax`
- For each element `ele1` in `array`: 
  - Track a `subsectionSum`
  - Iterate from `ele1` to the end, tracking `ele2`:
    - Add `ele2` to `subsectionSum`
    - If `subsectionSum` becomes negative -> break (move to next starting element)
- Return `currentMax`

### Efficiency

- All numbers positive -> No benefit (`N^2/2` iterations)
- Mix of pos/neg -> on average, `subsectionSum` will become neg. half of the time -> early breaks half of the time