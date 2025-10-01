package main

// 1518. Water Bottles
// There are numBottles water bottles that are initially full of water. You can exchange numExchange empty water bottles
// from the market with one full water bottle.
// The operation of drinking a full water bottle turns it into an empty bottle.
// Given the two integers numBottles and numExchange, return the maximum number of water bottles you can drink.

func numWaterBottles(numBottles int, numExchange int) int {
	drunk := numBottles
	var newBottles int
	for numBottles >= numExchange {
		newBottles = numBottles / numExchange
		numBottles = numBottles%numExchange + newBottles
		drunk += newBottles
	}

	return drunk
}
