package main

// 187. Repeated DNA Sequences
// The DNA sequence is composed of a series of nucleotides abbreviated as 'A', 'C', 'G', and 'T'.
//    For example, "ACGAATTCCG" is a DNA sequence.
// When studying DNA, it is useful to identify repeated sequences within the DNA.
// Given a string s that represents a DNA sequence, return all the 10-letter-long sequences (substrings) that occur more
// than once in a DNA molecule. You may return the answer in any order.

func findRepeatedDnaSequences(s string) []string {
	hash := make(map[string]int)
	limit := len(s) - 9
	for i := 0; i < limit; i++ {
		substr := s[i : i+10]
		if _, ok := hash[substr]; ok {
			hash[substr]++
		} else {
			hash[substr] = 1
		}
	}
	var result []string
	for h, c := range hash {
		if c > 1 {
			result = append(result, h)
		}
	}

	return result
}
