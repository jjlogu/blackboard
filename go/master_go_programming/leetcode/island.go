package main

// Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
// It's used mainly to print out to stdout
import "fmt"

func numIslands(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])
	var countIsland int

	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' && visited[i][j] == false {
				countIsland += 1
				bfs(grid, i, j, visited)
			}
		}
	}
	return countIsland
}
func bfs(grid [][]byte, r int, c int, visited [][]bool) {
	q := [][2]int{}

	rows := len(grid)
	cols := len(grid[0])

	q = append(q, [2]int{r, c})
	visited[r][c] = true
	fmt.Printf("start (%d, %d)\n", r, c)
	adjs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // up, down, left, right
	for len(q) > 0 {

		current := q[0]
		q = q[1:] // pop
		for _, adj := range adjs {
			i := current[0] + adj[0]
			j := current[1] + adj[1]
			if i >= 0 && i < rows && j >= 0 && j < cols && visited[i][j] == false && grid[i][j] == '1' {
				visited[i][j] = true
				fmt.Printf("(%d, %d) ", i, j)
				q = append(q, [2]int{i, j})
			}
		}
		fmt.Printf("\n")

	}
}
func dfs(grid [][]byte, r int, c int, visited [][]bool) {
	rows := len(grid)
	cols := len(grid[0])
	if r < 0 || r >= rows || c < 0 || c >= cols || visited[r][c] == true || grid[r][c] == '0' {
		fmt.Printf("\n")
		return
	}

	visited[r][c] = true
	fmt.Printf("(%d, %d) ", r, c)

	dfs(grid, r-1, c, visited)
	dfs(grid, r+1, c, visited)
	dfs(grid, r, c-1, visited)
	dfs(grid, r, c+1, visited)
}

func main() {
	grid := [][]byte{
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '1', '0'},
	}
	fmt.Printf("\nNumber of Island: %d\n", numIslands(grid))
}
