package main

import "fmt"

func main() {
	// Pointers
	// A pointer is a variable that stores the memory address of another variable.
	var x int = 42
	var p *int = &x // p is a pointer to x
	*p = 100        // change the value of x through the pointer
	x=200

	fmt.Println("x:", x) // Output: x: 100
	fmt.Println("p:", p, *p, &p) // Output: p: <address>

	// Pointers to pointers
	var pp **int = &p // pp is a pointer to p
	**pp = 200       // change the value of x through pp

	fmt.Println("x:", x) // Output: x: 200
}