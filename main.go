package main

import "github.com/Fefek12/tick/Server"

func main() {
	var game = Engine{grid: [3][3]string{
		{"x", "x", "h"},
		{"y", "y", "y"},
		{"z", "z", "z"},
	}}

	s := Server.NewServer(":8080")
	s.Start()
	game.grid[0][0] = "l"
	game.Render()
}
