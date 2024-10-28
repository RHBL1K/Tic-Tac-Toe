package gamefuncs

import "fmt"

func PrintBoard(board [3][3]string) {
	// for i := range board {
	// 	fmt.Println(" " + board[i][0] + " | " + board[i][1] + " | " + board[i][2])
	// 	if i < 2 {
	// 		fmt.Println("---+---+---")
	// 	}
	// }
	fmt.Println("    1   2   3")
    fmt.Println("  -------------")
    for i := 0; i < 3; i++ {
        fmt.Printf("%d |", i+1)
        for j := 0; j < 3; j++ {
            if board[i][j] == "" {
                fmt.Print("   ")
            } else {
                fmt.Printf(" %s ", board[i][j])
            }

            if j < 2 {
                fmt.Print("|")
            }
        }
        fmt.Println("|")
        if i < 2 {
            fmt.Println("  |---+---+---|")
        }
    }
    fmt.Println("  -------------")

}
