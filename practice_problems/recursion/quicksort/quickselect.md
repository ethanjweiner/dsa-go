# Quickselect

## Ideas

Given an array, find the `n`th lowest element in it

Partitioning: After a partition, all values left of the pivot are less than the
pivot, all values to the right are greater

We can scrap half of the values each time (no need to worry about sorting the half we don't care about).

Once the pivot index reaches the index we are looking for, we just return that value.

## Efficiency

There are still `O(log n)` layers of partitioning, but each layer only partitions a slice of halve as many elements as the previous.

The # of steps (approximate) for a slice of size `n` might look like:

`n + n/2 + n/4 + ... + 1` = `n + (n/2+n/4+..+)` = `n + n` = `2n` -> `O(n)`