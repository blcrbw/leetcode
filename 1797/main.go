package main

import "fmt"

// 1797. Design Authentication Manager
// There is an authentication system that works with authentication tokens. For each session, the user will receive a
//	 new authentication token that will expire timeToLive seconds after the currentTime. If the token is renewed, the
//	 expiry time will be extended to expire timeToLive seconds after the (potentially different) currentTime.
//
// Implement the AuthenticationManager class:
//
//    AuthenticationManager(int timeToLive) constructs the AuthenticationManager and sets the timeToLive.
//    generate(string tokenId, int currentTime) generates a new token with the given tokenId at the given currentTime in
//      seconds.
//    renew(string tokenId, int currentTime) renews the unexpired token with the given tokenId at the given currentTime
//      in seconds. If there are no unexpired tokens with the given tokenId, the request is ignored, and nothing happens
//    countUnexpiredTokens(int currentTime) returns the number of unexpired tokens at the given currentTime.
//
// Note that if a token expires at time t, and another action happens on time t (renew or countUnexpiredTokens), the
//   expiration takes place before the other actions.

type testArgs struct {
	token       string
	currentTime int
}

type AuthenticationManager struct {
	ttl    int
	tokens map[string]int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{ttl: timeToLive, tokens: make(map[string]int)}
}

func (a *AuthenticationManager) Generate(tokenId string, currentTime int) {
	a.tokens[tokenId] = currentTime + a.ttl

}

func (a *AuthenticationManager) Renew(tokenId string, currentTime int) {
	validUntil, ok := a.tokens[tokenId]
	if !ok {
		return
	} else if validUntil <= currentTime {
		delete(a.tokens, tokenId)
		return
	} else {
		a.tokens[tokenId] = currentTime + a.ttl
	}
}

func (a *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	count := 0
	for _, value := range a.tokens {
		if value > currentTime {
			count++
		}
	}
	return count
}

// "AuthenticationManager","renew","generate","countUnexpiredTokens","generate","renew","renew","countUnexpiredTokens"
func test(commands []string, args []testArgs, results []int) []int {

	if len(commands) != len(args) || len(commands) != len(results) {
		panic("Wrong number of arguments")
	}
	var lru AuthenticationManager
	var res int
	ret := make([]int, 0)
	for i := 0; i < len(commands); i++ {
		res = 0
		if commands[i] == "AuthenticationManager" {
			lru = Constructor(args[i].currentTime)
		} else if commands[i] == "renew" {
			lru.Renew(args[i].token, args[i].currentTime)
		} else if commands[i] == "generate" {
			lru.Generate(args[i].token, args[i].currentTime)
		} else if commands[i] == "countUnexpiredTokens" {
			res = lru.CountUnexpiredTokens(args[i].currentTime)
			fmt.Printf("GET result: expected %d, got %d\n", results[i], res)
		}
		ret = append(ret, res)
	}
	return ret
}
