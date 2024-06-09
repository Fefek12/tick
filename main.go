package main

func main() {
	var x = Engine{grid: [][]string{
		{"x", "x", "x"},
		{"y", "y", "y"},
		{"z", "z", "z"},
	}}
	x.Render()
}
