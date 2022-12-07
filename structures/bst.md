# BST

## Insertion

Given a `bst` and a `value`:
- If `value` < BST's `Data`:
  - If `Left` is `nil` -> assign `Left` to `value`
  - Otherwise, insert recursively on the left tree
- If `value` > BST's `Data`:
  - Same process for right
- Else, value must be equal (already in tree) -> do nothing