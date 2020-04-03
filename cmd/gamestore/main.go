package main

import (
	"bufio"
	"fmt"
	"os"

	c "github.com/stevedesilva/gofuncPointers/command"
	g "github.com/stevedesilva/gofuncPointers/game"
	f "github.com/stevedesilva/gofuncPointers/functask"
)

//     5- Refactor the runCmd() with the cmdXXX funcs.
//
//  Go back to main.go and change the existing code with
//  the new funcs that you've created. There are hints
//  inside the code.
//
// ---------------------------------------------------------

func main() {
	games := g.Load()
	byID := g.IndexByID(games)

	fmt.Printf("Inanc's Game store has %d Games.\n", len(games))

	in := bufio.NewScanner(os.Stdin)
	for {
		// menu()
		fmt.Printf(`
  > list   : lists all the Games
  > id N   : queries a Game by id
  > quit   : quits

`)

		if !in.Scan() || !c.RunCmd(in.Text(), games, byID) {
			break
		}
	}
}

func mainOld() {
	fmt.Println(f.Hello())

}
