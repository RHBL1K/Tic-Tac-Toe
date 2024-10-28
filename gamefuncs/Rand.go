package gamefuncs

import (
	"math/rand"
)

func TakeRand(board *[3][3]string, choice1 string, choice2 string) {
	for {
		// Генерируем случайные координаты строки и столбца
		randomRow := rand.Intn(len(board))
		randomCol := rand.Intn(len(board[0]))

		// Проверяем, пустая ли выбранная клетка
		if board[randomRow][randomCol] == " " {
			// Если пустая, ставим туда "O" и выходим из цикла
			board[randomRow][randomCol] = choice2
			break
		}
	}
}
