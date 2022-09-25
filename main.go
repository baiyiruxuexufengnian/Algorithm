package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

/*
Given a non-empty 2D array grid of 0's and 1's,
an island is a group of 1's (representing land)
connected 4-directionally(horizontal or vertical.)
You may assume all four edges of the grid are
surrounded by water.

Find the maximum area of an island in the
given 2D array.(If there is no island,
the maximum area is 0.)

Example 1:

[[0,0,1,0,0,0,0,1,0,0,0,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,1,1,0,1,0,0,0,0,0,0,0,0],
 [0,1,0,0,1,1,0,0,1,0,1,0,0],
 [0,1,0,0,1,1,0,0,1,1,1,0,0],
 [0,0,0,0,0,0,0,0,0,0,1,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,0,0,0,0,0,0,1,1,0,0,0,0]]
Given the above grid, return 6. Note the answer is not 11,
because the island must be connected 4-directionally.

Example 2:

[[0,0,0,0,0,0,0,0]]
Given the above grid, return 0.

Example 3:
[
 [1,0,0,0,1,0,1],
 [0,1,1,0,1,1,1],
 [0,0,0,0,0,0,1],
 [0,0,0,1,1,1,0],
 [0,0,0,1,1,0,0]
]
Given the above grid, return 6.

Note: The length of each dimension in the given grid does not exceed 50
*/

func wordBreak(s string, wordDict []string) bool {
	if len(wordDict) == 0 {
		return false
	}
	result := ""
	mem := make(map[string]bool)
	var wordBreakBacktrack func(wordDict []string) bool
	wordBreakBacktrack = func(wordDict []string) bool {
		if len(result) >= len(s) {
			if s == result {
				return true
			}
			return false
		}
		if s[:len(result)] != result {
			return false
		}
		if val, ok := mem[result]; ok {
			return val
		}
		for i := 0; i < len(wordDict); i ++ {
			result += wordDict[i]
			ans := wordBreakBacktrack(wordDict)
			if ans {
				mem[result] = true
				return true
			}
			mem[result] = false
			result = result[:len(result) - len(wordDict[i])]
		}
		return false
	}
	return wordBreakBacktrack(wordDict)
}

type TrieNode struct {
	end        bool
	cMap2Node  map[byte]*TrieNode
}

type Trie struct {
	size  int
	node  *TrieNode
}


func Constructor() Trie {
	return Trie{
		size: 0,
		node: &TrieNode{
			cMap2Node: make(map[byte]*TrieNode),
		},
	}
}

func (this *Trie) Insert(word string)  {
	bs := []byte(word)
	cur := this.node
	for i := 0; i < len(word); i ++ {
		if _, ok := cur.cMap2Node[bs[i]]; !ok {
			cur.cMap2Node[bs[i]] = &TrieNode{
				cMap2Node: map[byte]*TrieNode{},
			}
		}
		cur = cur.cMap2Node[bs[i]]
	}
	if !cur.end {
		cur.end = true
		this.size ++
	}
}


func (this *Trie) Search(word string) bool {
	bs := []byte(word)
	cur := this.node
	for _, c := range bs {
		if _, ok := cur.cMap2Node[c]; !ok {
			return false
		}
		cur = cur.cMap2Node[c]
	}
	return cur.end
}


func (this *Trie) StartsWith(prefix string) bool {
	bs := []byte(prefix)
	cur := this.node
	for _, c := range bs {
		log.Print("char: ", string(c))
		if _, ok := cur.cMap2Node[c]; !ok {
			return false
		}
		cur = cur.cMap2Node[c]
	}
	return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

//       0 1 2 3 4 5
// 输入: [3,2,1,5,6,4], k = 2
// 输出: 5
func findKthLargest(nums []int, k int) int {
	tp := len(nums) - k
	l := 0
	r := len(nums) - 1
	for l < r {
		pos := partition(nums, l, r)
		if tp == pos {
			return nums[pos]
		} else if pos < tp {
			l = pos + 1
		} else {
			r = pos - 1
		}
	}
	return l
}

func partition(nums []int, l, r int) int {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(r - l) + l
	pv := nums[index]
	nums[l], nums[index] = nums[index], nums[l]
	for l < r {
		for l < r && pv <= nums[r] { r -- }
		nums[r], nums[l] = nums[l], nums[r] // swap(nums[l], nums[r])
		for l < r && pv >= nums[l] { l ++ }
		nums[r], nums[l] = nums[l], nums[r] // swap(nums[l], nums[r])
	}
	return l
}

func quickSort(nums []int, l, r int) {
	if l < r {
		pos := partition(nums, l, r)
		quickSort(nums, l, pos - 1)
		quickSort(nums, pos + 1, r)
	}
}

func main() {
	nums := []int {3,2,3,1,2,4,5,5,6}
	quickSort(nums, 0, len(nums) - 1)
	fmt.Printf("nums: %v\n", nums)
	fmt.Printf("value: %v\n", findKthLargest(nums, 3))
}
