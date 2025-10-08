package leetcode

import "sort"

// 2300. Successful Pairs of Spells and Potions
// https://leetcode.com/problems/successful-pairs-of-spells-and-potions/
func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	result := make([]int, len(spells))

	for i, spell := range spells {
		minPotion := (success + int64(spell) - 1) / int64(spell)

		idx := sort.Search(len(potions), func(j int) bool {
			return int64(potions[j]) >= minPotion
		})

		result[i] = len(potions) - idx
	}

	return result
}
