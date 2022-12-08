# Greatest Number

Given an array of `numbers`

## O(N^2)

For each `number`:
- If `number` > every other number in `numbers`:
  - Return `number`

## O(N log N)

- Sort the `numbers` in ascending order
- Return the last

## O(N)

Iterate through `numbers`, reassigning a `greatestNumber` variable whenever a new greatest is found