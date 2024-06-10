package main

func main() {
	var x = Engine{grid: [][]string{
		{"x", "x", "x"},
		{"y", "y", "y"},
		{"z", "z", "z"},
	}}

	loading_screen("127.0.0.1")

	x.Render()

}
