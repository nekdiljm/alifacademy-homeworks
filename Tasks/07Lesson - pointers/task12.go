package main

import "fmt"

func main() {
	arr := []string{"Alice", "Carl", "Bob", "Carl", "Bob", "Bob", "Unknown"}
	var (
		aliceVotes int
		carlVotes  int
		bobVotes   int
	)

	for _, v := range arr {
		tallyVote(v, &aliceVotes, &carlVotes, &bobVotes)
	}
	fmt.Printf("Кандидат Alice: %v\n", aliceVotes)
	fmt.Printf("Кандидат Bob: %v\n", bobVotes)
	fmt.Printf("Кандидат Carl: %v\n", carlVotes)

}

func tallyVote(candidate string, countA, countB, countC *int) {
	switch candidate {
	case "Alice":
		*countA++
	case "Carl":
		*countB++
	case "Bob":
		*countC++
	}
}
