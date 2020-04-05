package command

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	g "github.com/stevedesilva/gofuncPointers/game"
)

// RunCmd func
func RunCmd(input string, games []g.Game, byID map[int]g.Game) bool {
	fmt.Println()

	cmd := strings.Fields(input)
	if len(cmd) == 0 {
		return true
	}

	switch cmd[0] {
	case "quit":
		return CmdQuit()

	case "list":
		return CmdList(games)

	case "id":
		return CmdByID(cmd, games, byID)

	case "save":
		return CmdSave(games)
	}

	return true

}

//
//  1- Create a new command in
//
//     For more information, see: "Encode" exercise from
//     the structs section.
//
//        Name  : cmdSave
//        Inputs: []game
//        Output: bool

// type jsonGame struct {
// 	ID    int    `json:"id,omitempty"`
// 	Name  string `json:"name,omitempty"`
// 	Genre string `json:"genre,omitempty"`
// 	Price int    `json:"price,omitempty"`
// }

// // NewJsonGame constructor
// func NewJsonGame(g.Game) *jsonGame {
// 	return
// }

// CmdSave func that encodes the
//     game store data to json and terminates the program.
//     Lastly, add it to runCmd func.
func CmdSave(games []g.Game) bool {

	jGames := make([]g.JSONGame, 0, len(games))

	for _, game := range games {
		jGames = append(jGames, g.NewJSONGame(game))
	}

	out, err := json.MarshalIndent(jGames, "", "\t")
	if err != nil {
		fmt.Println("sorry, can't save:", err)
		return true
	}
	fmt.Println(string(out))
	return false
}

//        This func returns true if it wants the program to
//        continue. When it returns false, the program will
//        terminate. So, all the commands that it calls need
//        to return true or false as well depending on whether
//        they want to cause the program to terminate or not.
//

//     2- Add a func that handles the quit command:
//
//        Name  : CmdQuit
//        Input : none
//        Output: bool
//

// CmdQuit func
func CmdQuit() bool {
	fmt.Println("bye!")
	return false
}

// CmdList func
//     3- Add a func that handles the list command:
//
//        Name  : CmdList
//        Inputs: []Game
//        Output: bool
//
//
func CmdList(games []g.Game) bool {
	// CmdList()
	for _, game := range games {
		g.PrintGame(game)
	}
	return true
}

// CmdByID func
//     4- Add a func that handles the id command:
//
//        Name  : CmdByID
//        Inputs: cmd []string, []Game, map[int]Game
//        Output: bool
//
func CmdByID(cmd []string, games []g.Game, byID map[int]g.Game) bool {
	// CmdByID
	if len(cmd) != 2 {
		fmt.Println("wrong id")
		return true
	}

	id, err := strconv.Atoi(cmd[1])
	if err != nil {
		fmt.Println("wrong id")
		return true
	}

	game, ok := byID[id]
	if !ok {
		fmt.Println("sorry. i don't have the Game")
		return true
	}

	g.PrintGame(game)

	return true
}
