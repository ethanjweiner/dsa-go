# Quicksort

## Efficiency

### Average & Best Case

- Time complexity of `Partition`:
  - Iterate until pointers meet or cross
  - Takes `O(n)` steps for pointers to meet/cross
- Number of `Partition` invocations:
  - Each `Quicksort` invocation invokes `Partition` once
  - Average case = (pivots end up somewhere in middle) -> `O(log N)` "layers" of partitions, with the total partitions for each layer performing `O(N)` steps
- `O(n log n)` time complexity in average case

### Worst Case

- No longer halving the array, just partitioning one less smaller each time
- -> ~ `n` `Partition` invocations: `N + (N - 1) + ... 1 + 2` total steps ~ `N^2 / 2` steps -> `O(n^2)`
- -> `O(n^2)` time complexity
