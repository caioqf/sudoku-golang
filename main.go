package main

import "fmt"

func generateGrid() [][]int {
	grid := make([][]int, 9)

	for i := range grid {
		grid[i] = make([]int, 9)
		for j := range grid[i] {
			grid[i][j] = 0
		}
	}

	return grid
}

func showGrid(grid [][]int) {
	for i, row := range grid {
		// itera os elementos na linha
		for j, element := range row {
			//verifica se é divisivel por 3 para printar a borda
			if (i*len(grid[0])+j)%3 == 0 {
				fmt.Printf("| ")
			}
			fmt.Printf("%d ", element)
		}
		// pula a linha
		fmt.Println()
	}
}

func main() {

	grid := generateGrid()
	showGrid(grid)

}
