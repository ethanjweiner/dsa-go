# Merge Two Sorted Lists

## Problem

- Input: Two sorted lists `list1` and `list2`
- Output: A new merged linked lists, sorted (return the head)

## Ideas

1, 2, 3, 4
2, 5, 4

1, 1, 2, 3,

- Iterative: Keep two pointers tracking the current value in each list, compare pointers -> add smaller
- Recursive: Since linked list is a recursive data structure, merge recursively

## Algorithm

Given `list1` and `list2`:
- If `list1` is empty -> return `list2`
- If `list2` is empty -> return `list1`
- If `list1`'s value <= `list2`'s value -> return `list1`'s head appended to the rest of the two lists merged
- If `list2`'s value <= `list1`'s value -> return `list2`'s head appended to the rest of the two lists merged