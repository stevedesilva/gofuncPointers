package main

import (
	f "github.com/stevedesilva/gofuncPointers/functask"

	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Refactor the game store to funcs - step #1
//
//  Remember the game store program from the structs exercises?
//  Now, it's time to refactor it to funcs.
//
//  Create games.go file
type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

//     2- Add a func that creates and returns a game.
func newGame(id, price int, name, genre string) game {
	item := item{id, name, price}
	return game{item, genre}
}

//     3- Add a func that adds a `game` to `[]game` and returns `[]game`:
func addGame(games []game, g game) []game {
	return append(games, g)
}

//     4- Add a func that uses newGame and addGame funcs to load up the game
//        store:
func load() (games []game) {
	games = addGame(games, newGame(1, 50, "god of war", "action adventure"))
	games = addGame(games, newGame(2, 40, "x-com 2", "strategy"))
	games = addGame(games, newGame(3, 20, "minecraft", "action sandbox"))
	return
}

//     5- Add a func that indexes games by ID:
func indexByID(games []game) map[int]game {
	byID := make(map[int]game)
	for _, g := range games {
		byID[g.id] = g
	}
	return
}

//     6- Add a func that prints the given game:
//
//        Name  : printGame
//        Inputs: game
//
//
//  Go back to main.go and change the existing code with
//  the new funcs that you've created. There are hints
//  inside the code.
//
// ---------------------------------------------------------
func printGame(g game) {
	fmt.Printf("#%d: %-15q %-20s $%d\n",
		g.id, g.name, "("+g.genre+")", g.price)
}

func main() {

	games := load()

	byID := indexByID(games)

	fmt.Printf("Inanc's game store has %d games.\n", len(games))

	in := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(`
  > list   : lists all the games
  > id N   : queries a game by id
  > quit   : quits
`)

		if !in.Scan() {
			break
		}

		fmt.Println()

		cmd := strings.Fields(in.Text())
		if len(cmd) == 0 {
			continue
		}

		switch cmd[0] {
		case "quit":
			fmt.Println("bye!")
			return

		case "list":
			for _, g := range games {
				printGame(g)
			}

		case "id":
			if len(cmd) != 2 {
				fmt.Println("wrong id")
				continue
			}

			id, err := strconv.Atoi(cmd[1])
			if err != nil {
				fmt.Println("wrong id")
				continue
			}

			g, ok := byID[id]
			if !ok {
				fmt.Println("sorry. i don't have the game")
				continue
			}

			printGame(g)
		}
	}
}

func mainOld() {
	fmt.Println(f.Hello())

}
