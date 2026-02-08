package leetcode

// 211. Design Add and Search Words Data Structure
// https://leetcode.com/problems/design-add-and-search-words-data-structure/description/

type WordDictionary struct {
	children [26]*WordDictionary
	isEnd    bool
}

func newWordDictionary() *WordDictionary {
	return &WordDictionary{
		children: [26]*WordDictionary{},
		isEnd:    false,
	}
}

func Constructor() WordDictionary {
	return *newWordDictionary()
}

func (wd *WordDictionary) AddWord(word string) {
	node := wd
	for _, c := range word {
		idx := c - 'a'
		if node.children[idx] == nil {
			node.children[idx] = newWordDictionary()
		}
		node = node.children[idx]
	}
	node.isEnd = true
}

func (wd *WordDictionary) Search(word string) bool {
	return searchWithDot(wd, word, 0)
}

func searchWithDot(wd *WordDictionary, word string, index int) bool {
	if index == len(word) {
		return wd.isEnd
	}
	if word[index] == '.' {
		for i := 0; i < 26; i++ {
			if wd.children[i] != nil && searchWithDot(wd.children[i], word, index+1) {
				return true
			}
		}
		return false
	}
	if wd.children[word[index]-'a'] == nil {
		return false
	}
	return searchWithDot(wd.children[word[index]-'a'], word, index+1)
}
