package main

import "strings"

// 1021. Remove Outermost Parentheses
//
// A valid parentheses string is either empty "", "(" + A + ")", or A + B, where A and B are valid parentheses strings,
// and + represents string concatenation.
//    For example, "", "()", "(())()", and "(()(()))" are all valid parentheses strings.
//
// A valid parentheses string s is primitive if it is nonempty, and there does not exist a way to split it into s = A + B,
// with A and B nonempty valid parentheses strings.
//
// Given a valid parentheses string s, consider its primitive decomposition: s = P1 + P2 + ... + Pk, where Pi are primitive
// valid parentheses strings.
//
// Return s after removing the outermost parentheses of every primitive string in the primitive decomposition of s.

func removeOuterParentheses(s string) string {
	c := 0
	var res []rune

	for _, i := range s {
		if i == '(' {
			if c > 0 {
				res = append(res, i)
			}
			c++
		} else {
			c--
			if c > 0 {
				res = append(res, i)
			}
		}
	}

	return string(res)
}

func removeOuterParentheses_byte(s string) string {
	c := 0
	res := []byte{}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			if c > 0 {
				res = append(res, s[i])
			}
			c++
		} else {
			c--
			if c > 0 {
				res = append(res, s[i])
			}
		}
	}

	return string(res)
}

func removeOuterParentheses_sb(s string) string {
	c := 0
	res := strings.Builder{}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			if c > 0 {
				res.WriteByte(s[i])
			}
			c++
		} else {
			c--
			if c > 0 {
				res.WriteByte(s[i])
			}
		}
	}

	return res.String()
}

func removeOuterParentheses_base(s string) string {
	primitives := []string{}
	c, p := 0, 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			c++
		} else {
			c--
		}
		if c == 0 {
			primitives = append(primitives, s[p:i+1])
			p = i + 1
		}
	}

	res := ""

	for _, j := range primitives {
		res += j[1 : len(j)-1]
	}

	return res
}
