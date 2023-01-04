/*
# PROBLEM

- Input: Two strings, `s` and `t`
- Output: Boolean indicating whether `s` is subsequence of `t`
- Definition of subsequence: `s` is subsequence of `t` if, after deleting some elements in `t`, `s == t` (relative order of `t` is preserved)

Edge cases: Duplicate chars in `t`
e.g. ahbbgdaddcbxxxc (fine, same result)

# IDEAS

TWO POINTERS (pointer for `s`, pointer for `t`)

Given strings `s` and `t`:
- Initialize `subIndex` to 0
- Iterate over `t`, tracking each `char`:
	- Once `char` equals `s[subIndex]`
		- Increment `subIndex`
- If `subIndex` == length(s), return true (order was found)

ahbgdc, abc
x       bc
xox     c
xoxoox  empty

-> true
TC: O(T)

*/