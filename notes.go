package main

import "fmt"

func main() {
	// Print function in "fmt"
	fmt.Println("Hello World.")

	// Variables
	var x int
	x = 1
	fmt.Printf("x=%v\n", x)
	var y float64
	y = 1.0
	fmt.Printf("y=%v\n", y)
	z := 1.0
	fmt.Printf("z is a %T\n", z)
	a, b := 1, 1.0
	fmt.Printf("a is a %T, while b is a %T\n", a, b)

	// if statement
	if y == z {
		fmt.Println("x is bigger than y")
	}

	// for loop
	for i := 0; i < 3; i++ {
		if i < 2 {
			continue
		}
		fmt.Println(i)
	}
}
