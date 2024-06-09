package main

import "fmt"

type Engine struct {
	grid [][]string
}

func (engine *Engine) Render() {
	for x := 0; x < len(engine.grid); x++ {
		fmt.Println("     |     |     ")
		for y := 0; y < len(engine.grid[x]); y++ {
			if y == 0 {
			} else if y == 1 {
				fmt.Printf("  |  %s  |  \n", engine.grid[x][y])
				// fmt.Printf("|  %s  |\n", engine.grid[x][y])
			} else {
				// fmt.Printf("  %s\n ", engine.grid[x][y])
			}
		}
		// fmt.Println("     |     |     d")
	}
}
