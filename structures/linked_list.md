# Linked List

## Deletion

Given a list `l` and a value to delete `value`:
- If the head node contains the value:
  - Reassign the head to the next node
  - Return
- If there is no next node, return
- If the next node contains the value:
  - Point to the current to the node next's next node
  - Return