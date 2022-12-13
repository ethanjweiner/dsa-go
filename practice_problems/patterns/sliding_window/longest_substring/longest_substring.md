# Longest Substring

## Problem

- Input: A `str`ing and positive integer `K`
- Output: Length of the longest substring w/ a maximum of `K` distinct characters

## Traditional Approach

For each character, keep trying to make the string one char bigger until `k` string has > `k` distinct characters (`O(n^2)` <= nested iteration)

## Sliding Window Approach

Idea: Keep incresing the window size, max window size = final result

Increase window size until a further increase would result in indistinct chars. 

araaci: window length becomes 4, then slide the window.

*Sub-problem*: Determine how many distinct chars a window has?
- `O(n)` approach: Iterate over substring, determine how many duplicates there are
- Additional idea: Maintain a hash tally of chars, increment/decrement char counts as needed.

## Algorithm

araaci (length = 6), k = 2

windowStart      | 0  0        0                                                        1
windowEnd        | 1  2        3                  4                  5                  6
longestSubLength | 0  0        0                                                        4
distintChars     | {} { a=>1 } { a => 1, r => 1 } { a => 2, r => 1 } { a => 3, r => 1 } { a => 3->2, r => 1, c => 1 }
currentLength    | 0  1        2                  3                  4                  5

*windowStart* represents index AT start of substring
*windowEnd* represents index AFTER the end of the substring

Given a `str` and `k`:
- Initialize `windowStart` to 0 and `windowEnd` to 1
- Initialize `longestSubstringLength` to `0`
- Initialize `distinctChars` to empty hash table
- Initialize `currentLength` to 0
- While `windowEnd` <= length of `str`: (`n` iterations)
  - Set `currentLength` to `windowEnd - windowStart`
  - Add char at `windowEnd - 1` to `distinctChars`
  - If length of `distintChars` > `k` and `currentLength - 1` > `longestSubstringLength`:
    - Set `longestSubstringLength` to `currentLength - 1`
    - Subtract the char at `windowStart` from `distinctChars`
      - If yields 0 -> delete char entirely
    - Increment `windowStart` by 1
  - Increment `windowEnd` by 1 
- [Final check] If `currentLength` > `longestSubstringLength`
  - Set `longestSubstringLength` to `currentLength - 1`
- Return `longestSubstringLength`

### Time Complexity

`O(n)`

