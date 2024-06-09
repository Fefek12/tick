package main

func main() {
	var game = Engine{grid: [][]string{
		{"x", "x", "h"},
		{"y", "y", "y"},
		{"z", "z", "z"},
	}}

	game.grid[0][0] = "l"
	game.Render()
}
