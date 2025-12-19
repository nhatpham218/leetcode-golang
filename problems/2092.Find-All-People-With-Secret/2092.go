package leetcode

import (
	"sort"
)

// 2092. Find All People With Secret
// https://leetcode.com/problems/find-all-people-with-secret/
//
// Solution: Sort meetings by time, then process meetings at the same time together.
// For each time slot, build a graph of connected people. If any person in that time
// slot knows the secret, propagate it to all connected people using DFS.
func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	knowsSecret := make([]bool, n)
	knowsSecret[0] = true
	knowsSecret[firstPerson] = true

	// Sort meetings by time to process them chronologically
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][2] < meetings[j][2]
	})

	// DFS to propagate secret to all connected people in the graph
	var propagateSecret func(person int, graph map[int][]int)
	propagateSecret = func(person int, graph map[int][]int) {
		for _, neighbor := range graph[person] {
			if !knowsSecret[neighbor] {
				knowsSecret[neighbor] = true
				propagateSecret(neighbor, graph)
			}
		}
	}

	// Process meetings grouped by time
	for i := 0; i < len(meetings); i++ {
		currentTime := meetings[i][2]

		// Build graph for all meetings at the current time
		graph := make(map[int][]int)
		hasSecretKeeper := false

		// Process all meetings at the same time
		j := i
		for j < len(meetings) && meetings[j][2] == currentTime {
			person1, person2 := meetings[j][0], meetings[j][1]

			// Check if anyone in this meeting already knows the secret
			if knowsSecret[person1] || knowsSecret[person2] {
				hasSecretKeeper = true
			}

			// Build bidirectional graph edges
			graph[person1] = append(graph[person1], person2)
			graph[person2] = append(graph[person2], person1)

			j++
		}

		// If someone in this time slot knows the secret, propagate it
		if hasSecretKeeper {
			for person := range graph {
				if knowsSecret[person] {
					propagateSecret(person, graph)
				}
			}
		}

		// Skip to the next time slot (j-1 because the outer loop will increment)
		i = j - 1
	}

	// Collect all people who know the secret
	result := make([]int, 0, n)
	for person := range knowsSecret {
		if knowsSecret[person] {
			result = append(result, person)
		}
	}
	return result
}
