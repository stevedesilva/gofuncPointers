// ---------------------------------------------------------
// EXERCISE: Fix the crash
//
// EXPECTED OUTPUT
//
//   brand: apple
// ---------------------------------------------------------

package main

import "fmt"

type computer struct {
	brand *string
}

func main() {
	var c *computer = &computer{}
    apple := "apple"
	fmt.Printf("apple: %s (%p)\n", apple,&apple)
	change(c, apple)
	fmt.Printf("brand: %s\n", *c.brand)
	fmt.Printf("apple.brand: %s (%p)\n", c,c.brand)
}

func change(c *computer, brand string) {
	fmt.Printf(">> apple: %s (%p)\n", brand,&brand)
	fmt.Printf(">>apple.brand: %s (%p)\n", c,c.brand)
	// auto def c
	// c.brand is a ptr to a val
	// point to another value
	c.brand = &brand
	fmt.Printf(">>apple.brand: %s (%p)\n", c,c.brand)
}