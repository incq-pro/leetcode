// 139. 单词拆分
// https://leetcode.cn/problems/word-break/

package word_break

type TrieNode struct {
	children map[uint8]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func newTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[uint8]*TrieNode),
		},
	}
}

func (t *Trie) insert(word string) {
	node := t.root
	for i := 0; i < len(word); i++ {
		ch := word[i]
		if next, ok := node.children[ch]; ok {
			node = next
		} else {
			next = &TrieNode{
				children: map[uint8]*TrieNode{},
			}
			node.children[ch] = next
			node = next
		}
	}
	node.isEnd = true
}

func (t *Trie) findAll(s string, start int) []int {
	var ret []int
	n := len(s)
	node := t.root
	for i := start; i < n; i++ {
		ch := s[i]
		if next, ok := node.children[ch]; ok {
			node = next
			if node.isEnd {
				ret = append(ret, i)
			}
		} else {
			break
		}
	}
	return ret
}

func wordBreak(s string, wordDict []string) bool {
	t := newTrie()
	for _, word := range wordDict {
		t.insert(word)
	}
	n := len(s)
	q := make([]bool, n)
	found := t.findAll(s, 0)
	if len(found) <= 0 {
		return false
	}
	for _, ff := range found {
		q[ff] = true
	}
	for k := 0; k < n; k++ {
		if !q[k] {
			continue
		}
		found = t.findAll(s, k+1)
		for _, ff := range found {
			q[ff] = true
		}
	}
	return q[n-1]
}

func WordBreak(s string, wordDict []string) bool {
	return wordBreak(s, wordDict)
}
