package main

import "fmt"

func main() {
	var (
		counter int
		V       int
		P       *int
	)

	// Section 1
	// P = &counter

	// if P == nil {
	// 	fmt.Printf("P is nil and its address is %p\n", P)
	// }

	// if P == &counter {
	// 	fmt.Printf("P is equal to the counters address: %p == %p\n", P, &counter)
	// }

	// Section 2
	counter = 100
	P = &counter
	V = *P

	//          counter : 100   addr: 0xc000014068
	fmt.Printf("counter : %-13d addr: %-13p\n ", counter, &counter)

	//          P       : 0xc000014068  addr: 0xc00000e020  *P: 100
	fmt.Printf("P		: %-13p addr: %-13p *P: %-13d\n ", P, &P, *P)

	fmt.Printf("V		: %-13d addr: %-13p\n ", V, &V)

	fmt.Println("\n...... change counter")

	counter = 10
	fmt.Printf("counter : %-13d addr: %-13p\n ", counter, &counter)

	fmt.Println("\n...... change counter in passVal")
	counter = passVal(counter)
	fmt.Printf("counter : %-13d addr: %-13p\n ", counter, &counter)

	fmt.Println("\n...... change counter in passPtrVal")
	passPtrVal(&counter) 
	fmt.Printf("counter : %-13d addr: %-13p\n ", counter, &counter)

}

func passPtrVal(pn *int) {
	fmt.Printf("pn		: %-13p addr: %-13p *pn: %-13d\n ", pn, &pn, *pn)
	*pn++ // same as (*pn)++
	fmt.Printf("pn		: %-13p addr: %-13p *pn: %-13d\n ", pn, &pn, *pn)

}

func passVal(n int) int {
	n = 50
	fmt.Printf("n : %-13d addr: %-13p\n ", n, &n)
	return n
}
