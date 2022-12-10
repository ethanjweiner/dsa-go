# Books with Authors

Given an array of `authors` (author id & name) & `books` (book id & title), produce an array that maps each book to the created author

## Efficiency

Let `a` be # of authors and `b` be # of books

1. Best imaginable Big O: `O(b)`: for each book, lookup associated author (in constant time)

Nested Iteration Approach = `O(b * a)`:
- For each book in `books`:
  - For each author in `authors`
    - If same `author_id` -> add to result array

`O(b + a)` approach:
- Transform `authors` into a hash table (key = `author_id`, value = `name`) (`a` steps)
- For each book in `books`: (`b` steps)
  - Lookup the name of the corresponding author ID (constant steps)
  - Add entry into array
