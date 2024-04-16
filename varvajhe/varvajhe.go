package varvajhe

import "fmt"

func countGoodSubstrings(S, P string) int {

	count := 0

	pFreq := make(map[byte]int)

	for i := 0; i < len(P); i++ {
		pFreq[P[i]]++
	}

	for i := 0; i < len(S)-len(P)+1; i++ {
		subFreq := make(map[byte]int)

		for j := 0; j < len(P); j++ {
			if S[i+j] != '?' {
				subFreq[S[i+j]]++
			}
		}

		isGood := true
		for k, v := range pFreq {
			if subFreq[k] != v {
				isGood = false
				break
			}
		}

		if isGood {
			count++
		}
	}

	return count
}

func main() {

	var S, P string
	fmt.Scan(&S, &P)

	fmt.Println(countGoodSubstrings(S, P))
}
