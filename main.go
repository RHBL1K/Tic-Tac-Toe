package main

import (
	"fmt"
	"MY-1st-project/gamefuncs"
)

func main() {
	// Создаю двумерный массив 3x3, пустую сетку
	
	board := [3][3]string{}

	// Добавляю пустые клетки пробелами
	for i := range board {
		for j := range board[i] {
			board[i][j] = " "
		}
	}

	gamefuncs.PrintBoard(board)
	Flag := false
	Flag2 := false
	Flagdraw := false
	
	var x, y int
	var choice1, choice2 string
	var gamerestart int


	// инпут хендлер на корректность ввода choice1 и ругаю Ерасыла
	for {
		fmt.Print("Введите X или O: \n")
		fmt.Scanln(&choice1)
	
		//Ругаю за капс
		if choice1 == "x" || choice1 == "o" {
			fmt.Println("Ты капс ставить не умеешь?")
			continue
		}
	
		//Ругаю за всё остальное
		if choice1 == "X" {
			choice2 = "O"
			break
		} else if choice1 == "O" {
			choice2 = "X"
			break
		} else {
			fmt.Println("Далбан, введи только Х или О, и без пробелов")
		}
	}



	for	{
		//спрашиваю координаты
		fmt.Printf("Введите координаты для %s (например, 1 1):\n", choice1)
		fmt.Scanln(&x, &y)
		for{
			// мини input handler, строковый ввод вместо целого числа не рассмотрен
			if(x>0 && x<4 && y>0 && y<4){
				break
			}else{
				fmt.Println("Далбаеб, у тебя клетки 3 на 3, куда ты лезешь за пределы")
				fmt.Printf("Введи координаты для %s (например, 1 1):\n", choice1)
				
				fmt.Scanln(&x, &y)
			}
		}
		//присваиваю значенения в массив и проверяю на победу
		if board[x-1][y-1] == " "{
			board[x-1][y-1] = choice1
			gamefuncs.Wincheck(&board,choice1,choice2,&Flag,&Flag2,&Flagdraw)
			if Flag == true{
				fmt.Println("Хорош, ты выйграл")
				break
			}
			if Flagdraw == true {
                fmt.Println("Бля, а как ты с ботом в ничью вышел?")
                break
            }
			gamefuncs.TakeRand(&board,choice1,choice2)
			gamefuncs.Wincheck(&board,choice1,choice2,&Flag,&Flag2,&Flagdraw)
			if Flag2 == true{
				fmt.Println("Ебать ты чмо боту проиграл")
				break
			}
			if Flagdraw == true {
                fmt.Println("Бля, а как ты с ботом в ничью вышел?")
                break
			}
		}else if board[x-1][y-1] == choice1 || board[x-1][y-1] == choice2{
			fmt.Println("Выбранная вами клетка уже заполнена, выбери другую")
		}
		
		gamefuncs.PrintBoard(board)
	}



	
	
	// проверка на повторный запуск
	gamefuncs.PrintBoard(board)
	fmt.Println("начать игру еще раз? 1 если да, 2 если нет")
	fmt.Scanln(&gamerestart)
	if gamerestart == 1{
		main()
	}else if gamerestart == 2{
		return
	}
}



