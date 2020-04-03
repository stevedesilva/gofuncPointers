package command

import (
	"fmt"
	"strconv"
	"strings"

	g "github.com/stevedesilva/gofuncPointers/game"
)

//     5- Refactor the runCmd() with the cmdXXX funcs.
//
//  Go back to main.go and change the existing code with
//  the new funcs that you've created. There are hints
//  inside the code.
//
// ---------------------------------------------------------

// func main() {
// 	games := f.Load()
// 	byID := f.IndexByID(Games)

// 	fmt.Printf("Inanc's Game store has %d Games.\n", len(Games))

// 	in := bufio.NewScanner(os.Stdin)
// 	for {
// 		// menu()
// 		fmt.Printf(`
//   > list   : lists all the Games
//   > id N   : queries a Game by id
//   > quit   : quits

// `)

// 		if !in.Scan() {
// 			break
// 		}

// 		// --- runCmd start ---
// 		continue := runCmd(in.Text(), games, byID)
// 		// --- runCmd end ---
// 	}
// }

// ---------------------------------------------------------
// EXERCISE: Refactor the Game store to funcs - step #2
//
//  Let's continue the refactoring from the previous
//  exercise. This time, you're going to refactor the
//  command handling logic.
//
//
//  Create commands.go file
//
//     1- Add a func that runs the given command from the user:
//
//        Name  : runCmd
//        Inputs: input string, []Game, map[int]Game
//        Output: bool
//

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
	}

	return true

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
