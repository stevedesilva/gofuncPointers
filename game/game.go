package game

import (
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Refactor the Game store to funcs - step #1
//
//  Remember the Game store program from the structs exercises?
//  Now, it's time to refactor it to funcs.
//
//  Create Games.go file

// Item struct
type Item struct {
	id    int
	name  string
	price int
}

// Game struct
type Game struct {
	Item
	genre string
}

// NewGame func
//     2- Add a func that creates and returns a Game.
func NewGame(id, price int, name, genre string) Game {
	Item := Item{id, name, price}
	return Game{Item, genre}
}

// AddGame func
//     3- Add a func that adds a `Game` to `[]Game` and returns `[]Game`:
func AddGame(Games []Game, g Game) []Game {
	return append(Games, g)
}

// Load struct
//     4- Add a func that uses NewGame and AddGame funcs to load up the Game
//        store:
func Load() (Games []Game) {
	Games = AddGame(Games, NewGame(1, 50, "god of war", "action adventure"))
	Games = AddGame(Games, NewGame(2, 40, "x-com 2", "strategy"))
	Games = AddGame(Games, NewGame(3, 20, "minecraft", "action sandbox"))
	return
}

// IndexByID func
//     5- Add a func that indexes Games by ID:
func IndexByID(Games []Game) map[int]Game {
	byID := make(map[int]Game)
	for _, g := range Games {
		byID[g.id] = g
	}
	return byID
}

// PrintGame func
//     6- Add a func that prints the given Game:
//
//        Name  : printGame
//        Inputs: Game
//
//
//  Go back to main.go and change the existing code with
//  the new funcs that you've created. There are hints
//  inside the code.
//
// ---------------------------------------------------------
func PrintGame(g Game) {
	fmt.Printf("#%d: %-15q %-20s $%d\n",
		g.id, g.name, "("+g.genre+")", g.price)
}
