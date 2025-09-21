package leetcode

import (
	"container/heap"
	"fmt"
)

// 1912. Design Movie Rental System
// https://leetcode.com/problems/design-movie-rental-system/

type IndexKey struct {
	Shop  int
	Movie int
}

func (k IndexKey) String() string {
	return fmt.Sprintf("%d_%d", k.Shop, k.Movie)
}

type MovieRentingSystem struct {
	rentedMovieHeap   *Movies
	unrentedMovies    map[int]*Movies
	rentedMovies      map[string]*RentingMovie
	indexingKeyOffset map[string]*RentingMovie
}

type RentingMovie struct {
	Shop     int
	Movie    int
	Price    int
	Offset   int
	IsRented bool
}

type Movies []*RentingMovie

func (m Movies) Len() int { return len(m) }

func (m Movies) Less(i int, j int) bool {
	if !m[i].IsRented {
		if m[i].Price == m[j].Price {
			return m[i].Shop < m[j].Shop
		}
	} else if m[i].Price == m[j].Price {
		if m[i].Shop == m[j].Shop {
			return m[i].Movie < m[j].Movie
		}
		return m[i].Shop < m[j].Shop
	}
	return m[i].Price < m[j].Price
}

func (m Movies) Swap(i int, j int) {
	m[i].Offset, m[j].Offset = j, i
	m[i], m[j] = m[j], m[i]
}

func (m *Movies) Push(x interface{}) {
	item := x.(*RentingMovie)
	item.Offset = m.Len()
	*m = append(*m, item)
}

func (m *Movies) Pop() interface{} {
	old := *m
	n := len(old)
	item := old[n-1]
	*m = old[0 : n-1]
	return item
}

func Constructor(n int, entries [][]int) MovieRentingSystem {
	unrentedMovies := make(map[int]*Movies)
	indexingMovies := make(map[string]*RentingMovie)

	for _, entry := range entries {
		if _, ok := unrentedMovies[entry[1]]; !ok {
			unrentedMovies[entry[1]] = &Movies{}
		}
		m := &RentingMovie{
			Shop:  entry[0],
			Movie: entry[1],
			Price: entry[2],
		}
		heap.Push(unrentedMovies[entry[1]], m)
		indexingMovies[IndexKey{
			Shop:  entry[0],
			Movie: entry[1],
		}.String()] = m
	}

	return MovieRentingSystem{
		indexingKeyOffset: indexingMovies,
		unrentedMovies:    unrentedMovies,
		rentedMovieHeap:   &Movies{},
		rentedMovies:      make(map[string]*RentingMovie),
	}
}

func (m *MovieRentingSystem) Search(movie int) []int {
	movies, ok := m.unrentedMovies[movie]
	if !ok {
		return nil
	}
	result := make([]int, 0, 5)
	eligibleMovies := make([]*RentingMovie, 0, 5)
	for movies.Len() > 0 && len(eligibleMovies) < 5 {
		item := heap.Pop(movies)
		movie := item.(*RentingMovie)
		eligibleMovies = append(eligibleMovies, movie)
		result = append(result, movie.Shop)
	}
	if len(eligibleMovies) == 0 {
		return nil
	}
	for _, e := range eligibleMovies {
		heap.Push(movies, e)
	}
	return result
}

func (m *MovieRentingSystem) Rent(shop int, movie int) {
	key := IndexKey{
		Shop:  shop,
		Movie: movie,
	}

	rentingMovie := m.indexingKeyOffset[key.String()]
	if rentingMovie == nil {
		return
	}

	heap.Remove(m.unrentedMovies[movie], rentingMovie.Offset)

	rentedMovie := &RentingMovie{
		Shop:     shop,
		Movie:    movie,
		Price:    rentingMovie.Price,
		IsRented: true,
	}
	m.rentedMovies[key.String()] = rentedMovie
	heap.Push(m.rentedMovieHeap, rentedMovie)
}

func (m *MovieRentingSystem) Drop(shop int, movie int) {
	key := IndexKey{
		Shop:  shop,
		Movie: movie,
	}
	rentedMovie, ok := m.rentedMovies[key.String()]
	if !ok {
		return
	}
	rentingMovie := &RentingMovie{
		Shop:  shop,
		Movie: movie,
		Price: rentedMovie.Price,
	}
	m.indexingKeyOffset[key.String()] = rentingMovie

	heap.Push(m.unrentedMovies[movie], rentingMovie)
	heap.Remove(m.rentedMovieHeap, rentedMovie.Offset)
	delete(m.rentedMovies, key.String())
}

func (m *MovieRentingSystem) Report() [][]int {
	result := make([][]int, 0, 5)
	historicals := make([]*RentingMovie, 0, 5)
	for m.rentedMovieHeap.Len() > 0 && len(historicals) < 5 {
		item := heap.Pop(m.rentedMovieHeap)
		movie := item.(*RentingMovie)
		historicals = append(historicals, movie)
		result = append(result, []int{movie.Shop, movie.Movie})
	}
	if len(historicals) == 0 {
		return nil
	}
	for _, item := range historicals {
		heap.Push(m.rentedMovieHeap, item)
	}
	return result
}
