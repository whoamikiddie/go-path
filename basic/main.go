package main

// --> importing a packages
import (
	"fmt" // -> this is for i/o print the statement
)

func main() {

	var a = "go"
	const b = 123
	c := 12333

	fmt.Print(a, b)
	fmt.Println(a, b)
	fmt.Printf("%v is a type of : %T\n", c, c)

}
