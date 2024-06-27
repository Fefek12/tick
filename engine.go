package main

import (
	"fmt"
)

type Engine struct {
	grid [3][3]string
}

func (engine *Engine) Render() {
	// clear()
	PrintTitle()
	for x := 0; x < len(engine.grid); x++ {
		if x == 0 {
			fmt.Println("     |     |     ")
		}
		for y := 0; y < len(engine.grid[x]); y++ {
			if y == 0 {
				fmt.Printf("  %s", engine.grid[x][y])
			} else if y == 1 {
				fmt.Printf("  |  %s  |  ", engine.grid[x][y])
				// fmt.Printf("|  %s  |\n", engine.grid[x][y])
			} else {
				fmt.Printf("%s\n", engine.grid[x][y])
			}
		}
		if x < 2 {
			fmt.Println("     |     |     \n-----|-----|-----\n     |     |     ")
		} else {
			fmt.Println("     |     |     ")
		}
	}
}
