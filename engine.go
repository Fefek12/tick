package main

import "fmt"

type Engine struct {
	grid [][]string
}

func (engine *Engine) Render() {
	for x := 0; x < len(engine.grid); x++ {
		for y := 0; y < len(engine.grid); y++ {
			fmt.Printf("%s", engine.grid[x][y])
		}
	}
}
