package main

import "fmt"

func main() {
	var states = map[string]int{
		"iowa":      8,
		"minnesota": 0,
		"arizona":   6,
		"florida":   5,
	}
	fmt.Println(rankStates(states))
}

func rankStates(states map[string]int) (stateRankings []string) {
	for state, rating := range states {
		if len(stateRankings) == 0 {
			stateRankings = append(stateRankings, state)
			continue
		}

		for i, rankedState := range stateRankings {
			if rating > states[rankedState] {
				stateRankings = append(stateRankings, "")
				copy(stateRankings[i+1:], stateRankings[i:])
				stateRankings[i] = state
				break
			}
			if i == len(stateRankings) {
				stateRankings = append(stateRankings, state)
			}
		}
	}
	return
}
