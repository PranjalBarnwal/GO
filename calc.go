package main

import "fmt"

func main() {
	fmt.Println("Cli Calculator")

	var decision string

	fmt.Println("Do you want to calculate 'y' for yes")
	fmt.Scan(&decision)

	for decision == "y" {

		fmt.Println("Enter two numbers")
		var n1, n2 int
		fmt.Scan(&n1)
		fmt.Scan(&n2)

		var operation string
		fmt.Println("Enter operation to perform (+,-,*,/)")
		fmt.Scan(&operation)

		switch operation {

		case "+":
			fmt.Printf("The result is %d\n", n1+n2)
		case "-":
			fmt.Printf("The result is %d\n", n1-n2)
		case "*":
			fmt.Printf("The result is %d\n", n1*n2)
		case "/":
			fmt.Printf("The result is %d\n", n1/n2)
		default:
			fmt.Println("Invalid Input")
		}

		fmt.Println("Do you want to calculate again 'y' for yes")
		fmt.Scan(&decision)

	}
}
