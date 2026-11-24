package main

import "testing"

func TestMinutesToRot(t *testing.T) {
	cases := []struct {
		in   [][]int
		want int
	}{
		{[][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}, 4},
		{[][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}, -1}, // bottom-left unreachable
		{[][]int{{0, 2}}, 0},                           // nothing fresh
		{[][]int{{1}}, -1},                             // fresh but no source
		{[][]int{{2, 2}, {1, 1}}, 1},                   // two sources spread together
		{[][]int{}, 0},
	}
	for _, c := range cases {
		if got := MinutesToRot(c.in); got != c.want {
			t.Errorf("MinutesToRot = %d, want %d", got, c.want)
		}
	}
}
