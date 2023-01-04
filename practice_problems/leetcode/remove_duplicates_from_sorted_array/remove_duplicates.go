package main

import "fmt"

/*
# PROBLEM

- Input: An array of `nums`
- Output: The same `nums`, w/ duplicates removed

# IDEAS

[0,0,1,1,1,2,2,3,3,4]

BRUTE FORCE

If two elements are the same, delete thee next & shift over
O(N^2) in worst case (iterate over every element, shift for every element)

Note: starting from the back won't work, b/c we need to shift all elements to
the front

OPTIMIZE

Anchor & runner

Questions to ask:
- Starting values?
- When to move?
- What do they do besides move?

[0,0,0,1,1,1,2,2,3,3,3,4]
A      R
[0,1,2,3,1,1,2,2,3,3,3,4]
   A     A             R

Read from runner
Write to anchor

Runner finds duplicates, anchor = next write

Edge Cases:

[0] -> Return [0]
[0, 0]
    A
    R

ALGORITHM

Given an array of `numbers`:
- If length is 1 -> return
- Initialize A & R to 1
- Loop:
	- While R == R - 1 AND R is not at end
		- Increment R
	- (R is no longer the same)
	- Write R To A
	- Move A to R

*/

func removeDuplicates(nums []int) {
	if len(nums) == 1 {
		return
	}

	a := 1

	for r := 1; r < len(nums); r++ {
		if nums[r] != nums[r-1] {
			nums[a] = nums[r]
			a++
		}
	}
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeDuplicates(nums)
	fmt.Println(nums) // [0, 1, 2, 3, 4, ....]

	nums2 := []int{0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeDuplicates(nums2)
	fmt.Println(nums2) // [0, 1, 2, 3, 4, ....]
}
