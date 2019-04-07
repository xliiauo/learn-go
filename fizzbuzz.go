package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {
		if i%15 == 0 {
			fmt.Println("fizz buzz")
			continue
		}

		if i%3 == 0 {
			fmt.Println("fizz")
			continue
		}

		if i%5 == 0 {
			fmt.Println("buzz")
			continue
		}

		fmt.Println(i)
	}
}
