package arrays

import (
	"math"
)

// BinarySearch_MedianOfTwoArrays - Given 2 sorted arrays this method needs to
// find the median of the 2 sorted array such that if the 2 arrays were merged
// and sorted there would be a common median. It needs to find it using binary
// search
func BinarySearch_MedianOfTwoArrays(a, b []int) float64 {
	/* Notes:
	The base binary algorithm goes as such:
		- low, hight = 0, len(a)
		- mid = low + (high - low) / 2
		- If a[mid] < target; hight = mid - 1
		- If a[mid] > target; low = mid + 1
		- If a[mid] == target; found!

	Using this concept, we can craft the following idea for find median b/w 2
	sorted arrays as if they were joined:

	If the arrays were joined, the median would split the joined into 2 parts:
		- left hand containing all smaller values
		- right hand containing all larger values

	If the total length were to be an Odd number - the median would be the largest
	value in the left hand array.

	If the totla length were to be an Even number - the median would be the smallest
	value in the right hand array.

	All we have to find is the partition between the 2 arrays.

	Here we can see that both the arrays will be partitioned into 2 groups:
		- The left hand of `a` would merge with left hand of `b`
		- The right hand of `b` would merge with right hand of `b`

	In the new groups created after the merger (let's call them `left` and `right`)
	we then simply hae to find the largest/smallest depending on the length.

	To create such a partition in both the arrays, we will do the following:
		1. Ensure len(a) < len(b); if not swap(a, b)
		2. Pick the mid value of a; say i
		3. We pick j which should be the partition in b such that on merger, both
		left and right will be of equal length; this is done by:
			- Split the total length by 2
			- Since we will add the left hand side of `a` to the left hand side of `b`
			- We will remove the length amount of a's left from b's left
			- Therefore we have: j = (len(a) + len(b)) / 2 - len(a[:i])
		4. Now we can say that both the arrays have been split into two:
			- al and ar = a[:i] and  a[i:]
			- bl and br = b[:j] and b[j:]
		5. And now we assume to have joined (bl + al) and (br + ar)
		6. For it to be the correct split:
			- max(al) < min(br) {note: here max and min are O(1) since they are the right
			and left hand sides of the array respectively}
			- min(ar) > max(bl)
		7. If max(al) > min(br), that means the value had to be in ar and there lies
		some other value in al that satisfies that condition. Therefore, we move the
		mid of a to the left (hight = mid)
		8. If min(ar) < max(bl), that means the value had to be in al and there lies
		some other value in ar that satisfies the condition. Therefore, we move the
		mid of a to the right (low = mid)
	*/

	// Array A is empty edge case
	if len(a) == 0 {
		m := len(b) / 2
		if len(b)%2 != 0 {
			return float64(b[m])
		}
		return (float64(b[m-1]) + float64(b[m])) / 2
	}

	// Array B is empty edge case
	if len(b) == 0 {
		m := len(a) / 2
		if len(a)%2 != 0 {
			return float64(a[m])
		}
		return (float64(a[m-1]) + float64(a[m])) / 2
	}

	s, l := a, b
	if len(s) > len(l) {
		s, l = b, a
	}

	ts := len(s) + len(l)
	low, high, i, j := 0, len(s)-1, 0, 0

	sL_max, sR_min := math.MaxInt, math.MinInt
	lL_max, lR_min := math.MaxInt, math.MinInt

	for sL_max > lR_min || lL_max > sR_min {
		i = low + (high-low)/2
		j = ts/2 - i

		sL_max = math.MinInt
		if i-1 >= 0 && i-1 < len(s) {
			sL_max = s[i-1]
		}
		sR_min = math.MaxInt
		if i >= 0 && i < len(s) {
			sR_min = s[i]
		}

		lL_max = math.MinInt
		if j-1 >= 0 && j-1 < len(l) {
			lL_max = l[j-1]
		}
		lR_min = math.MaxInt
		if j >= 0 && j < len(l) {
			lR_min = l[j]
		}

		if sL_max <= lR_min && lL_max <= sR_min {
			break
		}

		if sL_max > lR_min {
			high = i - 1
		} else {
			low = i + 1
		}
	}

	if ts%2 == 0 {
		return float64(max(sL_max, lL_max)+min(sR_min, lR_min)) / 2
	}
	return float64(min(sR_min, lR_min))
}
