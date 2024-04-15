package possiblesum

/*
You have a collection of coins, and you know the values of the coins and the quantity of each type of coin in it. You want to know how many distinct sums you can make from non-empty groupings of these coins.

Example

For coins = [10, 50, 100] and quantity = [1, 2, 1], the output should be
possibleSums(coins, quantity) = 9.

Here are all the possible sums:

50 = 50;
10 + 50 = 60;
50 + 100 = 150;
10 + 50 + 100 = 160;
50 + 50 = 100;
10 + 50 + 50 = 110;
50 + 50 + 100 = 200;
10 + 50 + 50 + 100 = 210;
10 = 10;
100 = 100;
10 + 100 = 110.
As you can see, there are 9 distinct sums that can be created from non-empty groupings of your coins.
*/

// https://booleanbit1.hashnode.dev/possible-sums
func solution(coins []int, quantity []int) int {
	// Create a HashSet to store all the possible sums
	sums := make(map[int]bool)

	// Add 0 as a possible sum, as we can always make a sum of 0
	sums[0] = true

	// Loop through all the coins
	for i := 0; i < len(coins); i++ {
		// Create a new HashSet to store the current set of sums
		currSums := make(map[int]bool)
		// Loop through all the previous possible sums
		for sum := range sums {
			// Loop through the quantity of the current coin
			for j := 1; j <= quantity[i]; j++ {
				// Add a new sum to the current set of sums
				// by adding the current coin multiplied by the quantity
				currSums[sum+j*coins[i]] = true
			}
		}

		// Add all the new sums to the main set of sums
		for key := range currSums {
			if _, found := sums[key]; !found {
				sums[key] = true
			}
		}
	}

	// Return the size of the set of sums, minus 1 (to exclude the 0 sum we added earlier)
	return len(sums) - 1
}

// https://github.com/mbaldini/challengequestions/blob/master/possibleSums.cs
func solution2(coins []int, quantity []int) int {
	// Create a HashSet to store all the possible sums
	sums := make(map[int]bool)

	// Add 0 as a possible sum, as we can always make a sum of 0
	sums[0] = true

	for i := 0; i < len(coins); i++ {
		sumsNow := []int{}
		for key := range sums {
			sumsNow = append(sumsNow, key)
		}
		for _, sum := range sumsNow {
			for j := 1; j <= quantity[i]; j++ {
				sums[sum+j*coins[i]] = true
			}
		}
	}

	return len(sums) - 1
}
