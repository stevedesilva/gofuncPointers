// ---------------------------------------------------------
// EXERCISE: Swap
//
//  Using funcs, swap values through pointers, and swap
//  the addresses of the pointers.
//
//
//  1- Swap the values through a func
//
//     a- Declare two float variables
//
//     b- Declare a func that can swap the variables' values
//        through their memory addresses
//
//     c- Pass the variables' addresses to the func
//
//     d- Print the variables
//
//
//  2- Swap the addresses of the float pointers through a func
//
//     a- Declare two float pointer variables and,
//        assign them the addresses of float variables
//
//     b- Declare a func that can swap the addresses
//        of two float pointers
//
//     c- Pass the pointer variables to the func
//
//     d- Print the addresses, and values that are
//        pointed by the pointer variables
//
// ---------------------------------------------------------

package main

import "fmt"

func main() {
	// exe1()
	exe2()

}

func exe1() {

	f1 := 1.0
	f2 := 2.0
	fmt.Printf(">> swapVal \n f1 %p (%v)\t f2 %p(%v)\n", &f1, f1, &f2, f2)
	swapVal(&f1, &f2)
	fmt.Printf("<< swapVal \n f1 %p (%v)\t f2 %p(%v)\n", &f1, f1, &f2, f2)

}

func swapVal(f1, f2 *float64) {
	// tmp := *f1
	// *f1 = *f2
	// *f2 = tmp
	*f1, *f2 = *f2, *f1
}

func exe2() {
	var (
		f1, f2 *float64
		v1     float64 = 1
		v2     float64 = 2
	)
	f1 = &v1
	f2 = &v2
	fmt.Printf(">> address swap \n f1 %p (%v)\t f2 %p(%v)\n", f1, *f1, f2, *f2)
	f1, f2 = swapAddress(f1, f2)
	fmt.Printf("<< address swap \n f1 %p (%v)\t f2 %p(%v)\n", f1, *f1, f2, *f2)
}

func swapAddress(f1, f2 *float64) (*float64, *float64) {

	return f2, f1
}
