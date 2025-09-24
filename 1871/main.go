package main

// 1871. Jump Game VII
// You are given a 0-indexed binary string s and two integers minJump and maxJump. In the beginning, you are standing at
// index 0, which is equal to '0'. You can move from index i to index j if the following conditions are fulfilled:
//    i + minJump <= j <= min(i + maxJump, s.length - 1), and
//    s[j] == '0'.
// Return true if you can reach index s.length - 1 in s, or false otherwise.

func canReach(s string, minJump int, maxJump int) bool {
	l := len(s)
	reachableSteps := make([]bool, l)
	reachableSteps[0] = s[0] == '0'

	// cnt is the number of steps from which we can jump into current one
	cnt := 0
	for i := minJump; i < l; i++ {
		if reachableSteps[i-minJump] {
			// Increase the counter if sliding window reached the available step
			cnt++
		}
		if i > maxJump && reachableSteps[i-maxJump-1] {
			// If sliding window moved from available step we have to decrease the counter
			cnt--
		}
		if s[i] == '0' && cnt > 0 {
			reachableSteps[i] = true
		}
	}

	return reachableSteps[l-1]
}
