package arrays

// SlidingWindow_LongestSubstringWithoutRepeatingChars - Solves the given problem
// statement using sliding window:
//
// Problem: Longest Substring Without Repeating Characters
// Given a string s, find the length of the longest substring without repeating
// characters.
//
// Constraints:
// 0 <= s.length <= 5 * 10⁴
//
// s consists of English letters, digits, symbols, and spaces.
//
// Example:
// Input: "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with a length of 3.
//
// Input: "bbbbb"
// Output: 1
//
// Input: "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with a length of 3.
func SlidingWindow_LongestSubstringWithoutRepeatingChars(main string) int {
	/** Notes:
	Here we will start of a small window and keep expanding it to the right till
	we encounter a new value that already exists in the window.

	Then we will shrink the window from the left hand side crossing the point
	where we encountered the first instance of the duplicate character.

	Before each shrink we will keep a note of the size of the window.

	And to make things faster we will keep the encountered character's indices
	into a map.

	Everytime we have to shrink a window - we will remove all the characters
	upto that duplicate character from the left most side of the window.
	*/

	memo := map[rune]int{}
	largest := 1
	counter := 0

	for index, char := range main {
		inIndex, isPresent := memo[char]

		// duplicate found
		if isPresent {
			// clear all prev charactes if duplicate found
			for key, value := range memo {
				if value <= inIndex {
					delete(memo, key)
					counter -= 1
				}
			}
		}

		memo[char] = index
		counter += 1

		if counter > largest {
			largest = counter
		}
	}

	return largest
}
