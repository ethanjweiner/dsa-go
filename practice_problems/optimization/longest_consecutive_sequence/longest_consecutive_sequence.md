# Longest Consecutive Sequencee

## O(N) Solution Ideas

- Find the min/max greedily (`O(N)`). Hash the sequence into a grouping of values (`O(N)`). Iterate from the min to the max, incrementing by 1. As you go, track the length of the current sequence. Once the sequence finishes, reassign its length to the max sequence length if that sequence is length is longer than the previous max.
- Find the min/max greedily (`O(N)`). Hash the sequence into a grouping of values (`O(N)`). Iterate over each value in `integers` again. For each value, incrementally look up consecutive values in the hash, accumulating a sequence length. If that sequence length is longer than the previous max, reassign.

Solution 2 is better because it avoids iterating over all integers betweewn a widely spread min/max.

Avoiding redundant steps: Only check integers which are the *bottom* number of the sequence (greedy approach)