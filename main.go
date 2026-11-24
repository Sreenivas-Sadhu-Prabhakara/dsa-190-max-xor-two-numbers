package main

import "fmt"

const topBit = 31

type bitTrie struct {
	children [2]*bitTrie
}

func (t *bitTrie) insert(v int) {
	n := t
	for b := topBit; b >= 0; b-- {
		bit := (v >> b) & 1
		if n.children[bit] == nil {
			n.children[bit] = &bitTrie{}
		}
		n = n.children[bit]
	}
}

// bestXor returns the maximum xor of v against any inserted value, walking
// the trie preferring the opposite bit at every level.
func (t *bitTrie) bestXor(v int) int {
	n := t
	result := 0
	for b := topBit; b >= 0; b-- {
		bit := (v >> b) & 1
		if n.children[1-bit] != nil {
			result |= 1 << b
			n = n.children[1-bit]
		} else {
			n = n.children[bit]
		}
	}
	return result
}

// MaxXor returns the maximum value of nums[i] XOR nums[j] over all pairs
// (i may equal j only when the array has one element).
func MaxXor(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	root := &bitTrie{}
	root.insert(nums[0])
	best := 0
	for _, v := range nums[1:] {
		if x := root.bestXor(v); x > best {
			best = x
		}
		root.insert(v)
	}
	return best
}

func main() {
	fmt.Println(MaxXor([]int{3, 10, 5, 25, 2, 8}))
}
