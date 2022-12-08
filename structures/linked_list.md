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

## Reverse (in place)

`l = Head -> A -> B -> C -> D`


### Ideas

- Starting from the end, delete from the end & prepend
  - Use a counter to determine when finished
- Starting from the beginning:

`l = Head -> A -> B -> C -> D`

Point Head to B and B to A, saving the next node B pointed to:
`l = Head -> B -> A ... C -> D` , saved node = `C`

Point Head to C and C to B, saved node = `D`
...

Continue until we've reached the final node (no longer points to another node)

List                         Node to Prepend   Node to Prepend Next
`Head -> A -> B ->  C -> D`  B (has a next)    C
                             Save `B`'s next 
`Head -> B -> A ... C -> D`  C (has a next)    D
                             Save `C`'s next
`Head -> C -> B -> A ... D`  D (has no next)   `nil`
                             Save `D`'s next
`Head -> D -> C -> B -> A`   `nil`
Done


## Idea 1

Given a list `l`:
- Save the current address of the `head` (first node)
- Loop:
  - Retrieve the `lastNode`
  - If the `lastNode` is the ori  ginal address of the first node:
    - Point that `lastNode` to `nil`
    - Break from the loop
  - Point the `lastNode` to the current node at the head
  - Point the head to the `lastNode`

## Idea 2

Head -> A -> B -> C -> D
            np   npn
Head -> B -> A ... C -> D
                   np   npn
Head -> C -> B -> A ... D     
                        np
Head -> D -> C -> B -> A  
np = `nil` -> done                     

Given a list `l`:
- If `head` has no next -> Return
- Initialize `nodeToPrepend` to the head's next
  - *Note*: Might just be able to assign to head itself -> no guard clause
- Until `nodeToPrepend` is `nil`:
  - Save the `nodeToPrepend`'s next into `nodeToPrependNext`
  - Point `nodeToPrepend` to the head node
  - Point the head node to `nodeToPrepend`
  - Assign `nodeToPrepend` to `nodeToPrependNext`