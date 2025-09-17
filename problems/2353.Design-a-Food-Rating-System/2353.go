package leetcode

import (
	"container/heap"
)

// 2353. Design a Food Rating System
// https://leetcode.com/problems/design-a-food-rating-system/

type Food struct {
	name   string
	rating int
}

// MaxHeap for foods (higher rating first, lexicographically smaller name for ties)
type FoodHeap []Food

func (h FoodHeap) Len() int { return len(h) }
func (h FoodHeap) Less(i, j int) bool {
	if h[i].rating == h[j].rating {
		return h[i].name < h[j].name // lexicographically smaller
	}
	return h[i].rating > h[j].rating // higher rating first
}
func (h FoodHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *FoodHeap) Push(x interface{}) {
	*h = append(*h, x.(Food))
}
func (h *FoodHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type FoodRatings struct {
	foodToCuisine map[string]string    // food -> cuisine
	foodToRating  map[string]int       // food -> rating
	cuisineHeaps  map[string]*FoodHeap // cuisine -> heap of foods
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	fr := FoodRatings{
		foodToCuisine: make(map[string]string),
		foodToRating:  make(map[string]int),
		cuisineHeaps:  make(map[string]*FoodHeap),
	}

	// Initialize the data structures
	for i, food := range foods {
		cuisine := cuisines[i]
		rating := ratings[i]

		// Map food to its cuisine and rating
		fr.foodToCuisine[food] = cuisine
		fr.foodToRating[food] = rating

		// Initialize heap for this cuisine if not exists
		if fr.cuisineHeaps[cuisine] == nil {
			heap := &FoodHeap{}
			fr.cuisineHeaps[cuisine] = heap
		}

		// Add food to the cuisine's heap
		heap.Push(fr.cuisineHeaps[cuisine], Food{name: food, rating: rating})
	}

	// Heapify all cuisine heaps
	for _, h := range fr.cuisineHeaps {
		heap.Init(h)
	}

	return fr
}

func (fr *FoodRatings) ChangeRating(food string, newRating int) {
	// Update the food's rating in the hash table
	fr.foodToRating[food] = newRating

	// Get the cuisine for this food
	cuisine := fr.foodToCuisine[food]

	// Add the updated food to the heap (we'll clean up stale entries in HighestRated)
	heap.Push(fr.cuisineHeaps[cuisine], Food{name: food, rating: newRating})
}

func (fr *FoodRatings) HighestRated(cuisine string) string {
	h := fr.cuisineHeaps[cuisine]

	// Clean up stale entries from the top of the heap
	for h.Len() > 0 {
		top := (*h)[0]
		// Check if the top entry has the current rating
		if fr.foodToRating[top.name] == top.rating {
			return top.name
		}
		// Remove stale entry
		heap.Pop(h)
	}

	return "" // Should never reach here based on problem constraints
}

/**
 * Your FoodRatings object will be instantiated and called as such:
 * obj := Constructor(foods, cuisines, ratings);
 * obj.ChangeRating(food,newRating);
 * param_2 := obj.HighestRated(cuisine);
 */
