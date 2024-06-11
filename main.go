package main

func main() {
	var game = Engine{grid: [3][3]string{
		{"x", "x", "h"},
		{"y", "y", "y"},
		{"z", "z", "z"},
	}}
	loading_screen("127.0.0.1")
	intro()
	// s := Server.NewServer(":8080")
	// s.Start()
	// game.grid[0][0] = "l"
	game.Render()
}
