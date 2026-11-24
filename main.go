package main

import "fmt"

// MinutesToRot returns how many minutes until every fresh orange (1) rots,
// when all rotten oranges (2) spread to 4-directional neighbors each minute.
// Returns -1 if some fresh orange can never be reached. Multi-source BFS:
// all initially rotten oranges start in the queue together.
func MinutesToRot(grid [][]int) int {
	type cell struct{ r, c int }
	queue := []cell{}
	fresh := 0
	for r, row := range grid {
		for c, v := range row {
			switch v {
			case 2:
				queue = append(queue, cell{r, c})
			case 1:
				fresh++
			}
		}
	}
	if fresh == 0 {
		return 0
	}
	minutes := 0
	dirs := []cell{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(queue) > 0 && fresh > 0 {
		minutes++
		next := []cell{}
		for _, cur := range queue {
			for _, d := range dirs {
				r, c := cur.r+d.r, cur.c+d.c
				if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) || grid[r][c] != 1 {
					continue
				}
				grid[r][c] = 2
				fresh--
				next = append(next, cell{r, c})
			}
		}
		queue = next
	}
	if fresh > 0 {
		return -1
	}
	return minutes
}

func main() {
	fmt.Println(MinutesToRot([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}))
}
