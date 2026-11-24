package main

import "testing"

func TestMaxXor(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{3, 10, 5, 25, 2, 8}, 28}, // 5 ^ 25
		{[]int{0}, 0},
		{[]int{2, 4}, 6},
		{[]int{8, 10, 2}, 10},
		{[]int{14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70}, 127},
		{[]int{}, 0},
		{[]int{7, 7}, 0}, // identical values xor to zero
	}
	for _, c := range cases {
		if got := MaxXor(c.in); got != c.want {
			t.Errorf("MaxXor(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}

// Property: agree with the O(n^2) scan.
func TestAgainstBrute(t *testing.T) {
	nums := []int{5, 29, 17, 3, 8, 42, 99, 61}
	want := 0
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x := nums[i] ^ nums[j]; x > want {
				want = x
			}
		}
	}
	if got := MaxXor(nums); got != want {
		t.Errorf("MaxXor = %d, brute = %d", got, want)
	}
}
