# Anagrams

## `O(N + M)` approach

- Tally counts of chars in each string (`N + M`) -> hash tables
- Return false if lengths of hash tables are different (constant)
  - Would have different # of chars
- For each count in str1: (`N`)
  - Check whether associated count for st2 (constant) is equal

`O(N + M + N)` = `O(2N + M)` = `O(N + M)`