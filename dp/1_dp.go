package dp

import "math"

// You are given an integer array coins representing coins of different
// denominations and an integer amount representing a total amount of money.
//
// Return the fewest number of coins that you need to make up that amount. If
// that amount of money cannot be made up by any combination of the coins, return -1.
//
// You may assume that you have an infinite number of each kind of coin.
//
// Example 1:
//
//	Input: coins = [1,2,5], amount = 11
//	Output: 3
//	Explanation: 11 = 5 + 5 + 1
//	Example 2:
//
// Constraints:
//
//	1 <= coins.length <= 12
//	1 <= coins[i] <= 231 - 1
//	0 <= amount <= 104
func DP_coinChange(coins []int, amount int) int {
	// DP table for number of coins required for each value from 0 to amount
	required := make([]int, amount+1)
	for i := range required {
		required[i] = math.MaxInt
	}
	// needs 0 coinst to take 0 amount
	required[0] = 0

	// for each value
	for v := 1; v < len(required); v++ {
		// we check if the given set of coins satisfy - the minimum possible coins
		// requried for that value
		for _, c := range coins {
			// if the value v required to be satisfied is less than the current coin's
			// value, then we skip trying to satisify the v with a bigger coin
			if c > v {
				continue
			}

			// otherwise

			// the current known minimum number of coins that is required to satisfy
			currentMinTaken := required[v]

			// the excess coin that would need to be satisfied if 1 of current coin
			// is to be taken
			excess := v - c

			// the known minimum number of coins that is required to satisfy the excess
			minTakeForExcess := required[excess]

			// if the excess wasn't satisfied - we cannot satisfy this value with this coin
			if minTakeForExcess == math.MaxInt {
				continue
			}

			// +1 is of the coin we are taking itself
			required[v] = min(currentMinTaken, minTakeForExcess+1)
		}
	}

	if required[amount] == math.MaxInt {
		// amount was never satisfied with any amount taken
		return -1
	}
	return required[amount]
}
