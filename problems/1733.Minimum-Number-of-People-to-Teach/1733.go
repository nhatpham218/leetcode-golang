package leetcode

// 1733. Minimum Number of People to Teach
// On a social network consisting of m users and some friendships between users, two users can communicate with each other if they know a common language.
//
// You are given an integer n, an array languages, and an array friendships where:
// There are n languages numbered 1 through n,
// languages[i] is the set of languages the i​​​​​​th​​​​ user knows, and
// friendships[i] = [u​​​​​​i​​​, v​​​​​​i] denotes a friendship between the users u​​​​​​​​​​​i​​​​​ and vi.
//
// You can choose one language and teach it to some users so that all friends can communicate with each other.
// Return the minimum number of users you need to teach.
//
// Note that friendships are not transitive, meaning if x is a friend of y and y is a friend of z, this doesn't guarantee that x is a friend of z.
func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	m := len(languages)

	// Convert languages to sets for faster lookup
	langSet := make([]map[int]bool, m)
	for i := 0; i < m; i++ {
		langSet[i] = make(map[int]bool)
		for _, lang := range languages[i] {
			langSet[i][lang] = true
		}
	}

	// Check if two users can communicate (have common language)
	canCommunicate := func(u, v int) bool {
		for lang := range langSet[u] {
			if langSet[v][lang] {
				return true
			}
		}
		return false
	}

	// Find all problematic friendships (can't communicate)
	var problemPairs [][]int
	for _, friendship := range friendships {
		u, v := friendship[0]-1, friendship[1]-1 // Convert to 0-indexed
		if !canCommunicate(u, v) {
			problemPairs = append(problemPairs, []int{u, v})
		}
	}

	// If no problematic pairs, no teaching needed
	if len(problemPairs) == 0 {
		return 0
	}

	minTeach := m // Worst case: teach everyone

	// Try each language
	for lang := 1; lang <= n; lang++ {
		// Find users involved in problematic pairs who don't know this language
		needToLearn := make(map[int]bool)

		for _, pair := range problemPairs {
			u, v := pair[0], pair[1]

			// If neither knows the language, both need to learn
			uKnows := langSet[u][lang]
			vKnows := langSet[v][lang]

			if !uKnows && !vKnows {
				needToLearn[u] = true
				needToLearn[v] = true
			} else if !uKnows {
				needToLearn[u] = true
			} else if !vKnows {
				needToLearn[v] = true
			}
		}

		if len(needToLearn) < minTeach {
			minTeach = len(needToLearn)
		}
	}

	return minTeach
}
