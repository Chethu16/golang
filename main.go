package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println("Enter the number of a and b and c")
	fmt.Scanln(&a, &b, &c)
	fmt.Println("The addition of a and b is", a-b+c)

}
