package leetcode

import (
	"sort"
	"strconv"
	"strings"
)

// 3433. Count Mentions Per User
// https://leetcode.com/problems/count-mentions-per-user/description/
func countMentions(numberOfUsers int, events [][]string) []int {
	mentions := make([]int, numberOfUsers)
	// offlineUntil[i] = timestamp when user i comes back online (0 means online)
	offlineUntil := make([]int, numberOfUsers)

	// Sort events by timestamp, with OFFLINE events processed before MESSAGE at same timestamp
	sort.Slice(events, func(i, j int) bool {
		tsI, _ := strconv.Atoi(events[i][1])
		tsJ, _ := strconv.Atoi(events[j][1])
		if tsI == tsJ {
			return events[i][0] == "OFFLINE" && events[j][0] != "OFFLINE"
		}
		return tsI < tsJ
	})

	for _, event := range events {
		timestamp, _ := strconv.Atoi(event[1])
		eventType := event[0]

		if eventType == "OFFLINE" {
			userID, _ := strconv.Atoi(event[2])
			offlineUntil[userID] = timestamp + 60
		} else {
			// Process MESSAGE event
			mentionsStr := event[2]

			switch mentionsStr {
			case "HERE":
				// Mention all online users
				for i := 0; i < numberOfUsers; i++ {
					if timestamp >= offlineUntil[i] {
						mentions[i]++
					}
				}
			case "ALL":
				// Mention all users (including offline)
				for i := 0; i < numberOfUsers; i++ {
					mentions[i]++
				}
			default:
				// Parse individual user IDs (format: "id0 id1 id2")
				tokens := strings.Split(mentionsStr, " ")
				for _, token := range tokens {
					if strings.HasPrefix(token, "id") {
						if userID, err := strconv.Atoi(token[2:]); err == nil && userID < numberOfUsers {
							mentions[userID]++
						}
					}
				}
			}
		}
	}

	return mentions
}
