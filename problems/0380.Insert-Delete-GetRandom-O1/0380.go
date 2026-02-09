package leetcode

import "math/rand"

// 380. Insert Delete GetRandom O(1)
// https://leetcode.com/problems/insert-delete-getrandom-o1/description/
type RandomizedSet struct {
	valToIndex map[int]int // maps value to its index in the slice
	values     []int       // stores all values for random access
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		valToIndex: make(map[int]int),
		values:     make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	// Check if value already exists
	if _, exists := this.valToIndex[val]; exists {
		return false
	}
	// Add to slice and map
	this.valToIndex[val] = len(this.values)
	this.values = append(this.values, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	// Check if value exists
	index, exists := this.valToIndex[val]
	if !exists {
		return false
	}
	
	// Swap with last element
	lastIndex := len(this.values) - 1
	lastVal := this.values[lastIndex]
	this.values[index] = lastVal
	this.valToIndex[lastVal] = index
	
	// Remove last element
	this.values = this.values[:lastIndex]
	delete(this.valToIndex, val)
	
	return true
}

func (this *RandomizedSet) GetRandom() int {
	// Pick a random index from the slice
	randomIndex := rand.Intn(len(this.values))
	return this.values[randomIndex]
}
