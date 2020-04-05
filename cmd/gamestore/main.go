package main

import (
	"bufio"
	"fmt"
	"os"

	c "github.com/stevedesilva/gofuncPointers/command"
	f "github.com/stevedesilva/gofuncPointers/functask"
	g "github.com/stevedesilva/gofuncPointers/game"
)

// ---------------------------------------------------------
// EXERCISE: Refactor the game store to funcs - step #3
//
//  Let's continue from the previous exercise. This time,
//  you're going to add json encoding and decoding using
//  funcs.
//
//  1- Create a new command in a func that encodes the
//     game store data to json and terminates the program.
//     Lastly, add it to runCmd func.
//
//     For more information, see: "Encode" exercise from
//     the structs section.
//
//        Name  : cmdSave
//        Inputs: []game
//        Output: bool
//
//  2- Refactor the load() to load the games from the
//     `data` constant (it's in the games.go as well).
//
//     For more information, see: "Decode" exercise from
//     the structs section.
//
// ---------------------------------------------------------
func main() {
	games := g.Load()
	byID := g.IndexByID(games)

	fmt.Printf("Inanc's Game store has %d Games.\n", len(games))

	in := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(`
  > list   : lists all the Games
  > id N   : queries a Game by id
  > quit   : quits
  > save   : exports the data to json and quits
`)

		if !in.Scan() || !c.RunCmd(in.Text(), games, byID) {
			break
		}
	}
}

func mainOld() {
	fmt.Println(f.Hello())

}
