package main

import "fmt"

func main() {
	// fmt.Println("Even Odd Checker")

	// fmt.Print("Enter your choice number: ")
	// var n int
	// fmt.Scan(&n)

	// if n%2 == 0 {
	// 	fmt.Printf("You entered an even number: %d\n", n)
	// } else {
	// 	fmt.Printf("You entered an odd number: %d\n", n)
	// }

	fmt.Println("Triangle classifier")

	fmt.Print("Enter three float numbers")
	var n1, n2, n3 float64
	fmt.Println("Enter the length of first side")
	fmt.Scan(&n1)
	fmt.Println("Enter the length of second side")
	fmt.Scan(&n2)
	fmt.Println("Enter the length of third side")
	fmt.Scan(&n3)

	if n1+n2 < n3 || n2+n3 < n1 || n1+n3 < n2 {
		fmt.Println("Not a valid triangle")
	} else {
		fmt.Println("Valid Triangle")
	}
}
