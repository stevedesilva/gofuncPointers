package main

import "fmt"

type computer struct {
	brand string
}

func main() {
	// create a nil pointer with the type of pointer to a computer
	var c *computer
	// compare the pointer variable to a nil value, and say it's nil
	if c == nil {
		fmt.Println("Computer pointer is nil")
	}
	// create an apple computer by putting its address to a pointer variable
	apple := &computer{brand: "apple"}
	// put the apple into a new pointer variable
	c = apple
	// compare the apples: if they are equal say so and print their addresses
	if apple == c {
		fmt.Printf("Apple %p(%v) are equal to c %p(%v) \n", apple, *apple, c, *c)
	}
	// create a sony computer by putting its address to a new pointer variable
	sony := &computer{brand: "sony"}

	// compare apple to sony, if they are not equal say so and print their
	// addresses
	if apple != sony {
		fmt.Printf("Apple %p(%v) are not equal to c %p(%v) \n", apple, *apple, sony, *sony)
	}
	// put apple's value into a new ordinary variable
	appleVal := *apple

	// print apple pointer variable's address, and the address it points to
	// and, print the new variable's addresses as well
	fmt.Printf("Apple %p (%v), appleVal %p(%v) \n", &apple, apple, &appleVal, appleVal)

	// compare the value that is pointed by the apple and the new variable
	// if they are equal say so
	if appleVal == *apple {
		fmt.Printf("Apple values are equal\n")
	}

	// print the values:
	// the value that is pointed by the apple and the new variable
	fmt.Printf("Apple (%v), appleVal (%v) \n", *apple, appleVal)

	// create a new function: change
	// the func can change the given computer's brand to another brand

	// change sony's brand to hp using the func — print sony's brand
	change(sony,"hp")
	fmt.Printf("Sony (%v) \n", sony)

	// write a func that returns the value that is pointed by the given *computer
	// print the returned value
	fmt.Printf("appleVal                  : %+v\n", valueOf(apple))

	// write a new constructor func that returns a pointer to a computer
	// and call the func 3 times and print the returned values' addresses
	
	fmt.Printf("Value (%v)\n",newFunc("chromebook") )
	fmt.Printf("Value (%v)\n",newFunc("dell") )
	fmt.Printf("Value (%v)\n",newFunc("asus") )
}

func newFunc(value string) *computer {
	return &computer{brand: value}
}

func change(from *computer, brand string) {
	from.brand = brand
}

func valueOf(from *computer) computer {
	return *from
}
