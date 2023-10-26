package main

import "fmt"

func sum(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func mul(a, b int) int {
	return a * b
}
func div(a, b float64) float64 {
	return a / b
}
func display(msg string) {
	fmt.Println(msg)
}
func main() {
	fmt.Println("First Commit")
	fmt.Println(sum(1, 2))
	fmt.Println(sub(1, 2))
	fmt.Println(mul(1, 2))
	fmt.Println(div(1, 2))
	display("Hello")
}
