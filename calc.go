package main

import "fmt"

func main() {

	var num1 float32
	var num2 float32
	var result float32
	var operator string
	

		fmt.Println("Enter first no")
		fmt.Scan(&num1)
		fmt.Println("Enter second no")
		fmt.Scan(&num2)
		fmt.Printf("enter the operator: +,*,-,/\n")
		fmt.Scan(&operator)

		switch operator {

		case "+":
			result = num1 + num2
			fmt.Printf("%v + %v is: %v\n", num1, num2, result)

		case "-":
			result = num1 - num2
			fmt.Printf("%v - %v is: %v\n", num1, num2, result)

		case "/":
			result = num1 / num2
			fmt.Printf("%v / %v is: %.5v\n", num1, num2, result)
 
		case "*":
			result = num1 * num2
			fmt.Printf("%v * %v is: %v\n", num1, num2, result)

		default:
			fmt.Printf("invalid operator type\n")
		}

	
}