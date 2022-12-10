# Temperature Readings

Given an array of `temperature readings`, sort in ascending order in `O(N)` time.

## Ideas

Given an array or `temps`:
- Hash the `temps` into a hash table (as a `tally`) = `N` steps
- Initialize a `sortedTemps` to an empty array
- Iterate from 97.0 to 99.0, in increments of +0.1 = constant # of steps
  - Lookup the value in the `tally` + add to `sortedTemps` based on # of times in `tally` = `N` total additions
- Return the new `sortedTemps`

*Multiply by ten to avoid working in floating-point increments*

Overall time complexity: `O(N)`