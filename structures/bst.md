# BST

## Insertion

Given a `bst` and a `value`:
- If `value` < BST's `Data`:
  - If `Left` is `nil` -> assign `Left` to `value`
  - Otherwise, insert recursively on the left tree
- If `value` > BST's `Data`:
  - Same process for right
- Else, value must be equal (already in tree) -> do nothing

## Deletion

Given a `bst`, `value`
- If `value` < `bst.Data`
  - Delete from left
- If `value` > `bst.Data`
  - Delete from right
- Else, `value` = `bst.Data`:
  - If there is only one child, then make that child the successor
  - Search for the `successor` of `bst`, re-inserting child nodes as needed

### Search for the successor (w/ two children)

Given a `bst`:
- Find the smallest node in the right side of the BST
  - Continually search the left until the left is `nil`
- Once the successor is found, remove the pointer to that successor

## In-Order Traversal

Given a `bst` and a function `f`:
- If `bst` is nil -> Return (done traversing)
- If the `bst`'s left is not `nil`:
  - Traverse the left of the `bst`
  - Call `f` on the `bst.Data`
- Else (left has no data, good to go)
  - Call `f` on the `bst.Data`
  - Traverse the right of the `bst`
  

## Efficiency

### Searching

Given a balanced BST w/ `n` elements:
- Discard half elements after comparison
- `x` = # of comparisons: `n * (1/2)^x = 1` -> `2^x = n` -> `x = log n`
- Therefore searching is `O(log n)`
